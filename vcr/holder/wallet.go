/*
 * Copyright (C) 2022 Nuts community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package holder

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
	ssi "github.com/nuts-foundation/go-did"
	"github.com/nuts-foundation/go-did/did"
	"github.com/nuts-foundation/go-did/vc"
	"github.com/nuts-foundation/go-stoabs"
	"github.com/nuts-foundation/nuts-node/auth/oauth"
	"github.com/nuts-foundation/nuts-node/core"
	"github.com/nuts-foundation/nuts-node/crypto"
	"github.com/nuts-foundation/nuts-node/jsonld"
	"github.com/nuts-foundation/nuts-node/vcr/credential"
	"github.com/nuts-foundation/nuts-node/vcr/log"
	"github.com/nuts-foundation/nuts-node/vcr/pe"
	"github.com/nuts-foundation/nuts-node/vcr/signature"
	"github.com/nuts-foundation/nuts-node/vcr/signature/proof"
	"github.com/nuts-foundation/nuts-node/vcr/verifier"
	"github.com/nuts-foundation/nuts-node/vdr/resolver"
	"strings"
	"time"
)

const statsShelf = "stats"

// ErrNoCredentials is returned when no matching credentials are found in the wallet based on a PresentationDefinition
var ErrNoCredentials = errors.New("no matching credentials")

var credentialCountStatsKey = stoabs.BytesKey("credential_count")

type wallet struct {
	keyResolver   resolver.KeyResolver
	keyStore      crypto.KeyStore
	verifier      verifier.Verifier
	jsonldManager jsonld.JSONLD
	walletStore   stoabs.KVStore
}

// New creates a new Wallet.
func New(
	keyResolver resolver.KeyResolver, keyStore crypto.KeyStore, verifier verifier.Verifier, jsonldManager jsonld.JSONLD,
	walletStore stoabs.KVStore) Wallet {
	return &wallet{
		keyResolver:   keyResolver,
		keyStore:      keyStore,
		verifier:      verifier,
		jsonldManager: jsonldManager,
		walletStore:   walletStore,
	}
}

// BuildParams contains the parameters that will be included in the signature of the verifiable presentation
type BuildParams struct {
	Audience string
	Expires  time.Time
	Nonce    string
}

func (h wallet) BuildSubmission(ctx context.Context, walletDID did.DID, presentationDefinition pe.PresentationDefinition, acceptedFormats map[string]map[string][]string, params BuildParams) (*vc.VerifiablePresentation, *pe.PresentationSubmission, error) {
	// get VCs from own wallet
	credentials, err := h.List(ctx, walletDID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve wallet credentials: %w", err)
	}

	// match against the wallet's credentials
	// if there's a match, create a VP and call the token endpoint
	// If the token endpoint succeeds, return the access token
	// If no presentation definition matches, return a 412 "no matching credentials" error
	builder := presentationDefinition.PresentationSubmissionBuilder()
	builder.AddWallet(walletDID, credentials)

	// Find supported VP format, matching support from:
	// - what the local Nuts node supports
	// - the presentation definition "claimed format designation" (optional)
	// - the verifier's metadata (optional)
	formatCandidates := credential.OpenIDSupportedFormats(oauth.DefaultOpenIDSupportedFormats())
	formatCandidates = formatCandidates.Match(credential.OpenIDSupportedFormats(acceptedFormats))
	if presentationDefinition.Format != nil {
		formatCandidates = formatCandidates.Match(credential.DIFClaimFormats(*presentationDefinition.Format))
	}
	// todo: next to the format selection, also check for algorithm support
	format := pe.ChooseVPFormat(formatCandidates.Map)
	if format == "" {
		return nil, nil, errors.New("requester, verifier (authorization server metadata) and presentation definition don't share a supported VP format")
	}
	presentationSubmission, signInstructions, err := builder.Build(format)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build presentation submission: %w", err)
	}
	if signInstructions.Empty() {
		return nil, nil, ErrNoCredentials
	}

	// todo: support multiple wallets
	vp, err := h.BuildPresentation(ctx, signInstructions[0].VerifiableCredentials, PresentationOptions{
		Format: format,
		ProofOptions: proof.ProofOptions{
			Created:   time.Now(),
			Challenge: &params.Nonce,
			Domain:    &params.Audience,
			Expires:   &params.Expires,
			Nonce:     &params.Nonce,
		},
	}, &walletDID, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create verifiable presentation: %w", err)
	}
	return vp, &presentationSubmission, nil
}

func (h wallet) BuildPresentation(ctx context.Context, credentials []vc.VerifiableCredential, options PresentationOptions, signerDID *did.DID, validateVC bool) (*vc.VerifiablePresentation, error) {
	var err error
	if signerDID == nil {
		signerDID, err = credential.ResolveSubjectDID(credentials...)
		if err != nil {
			return nil, fmt.Errorf("unable to resolve signer DID from VCs for creating VP: %w", err)
		}
	}

	kid, _, err := h.keyResolver.ResolveKey(*signerDID, nil, resolver.NutsSigningKeyType)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve assertion key for signing VP (did=%s): %w", *signerDID, err)
	}
	key, err := h.keyStore.Resolve(ctx, kid.String())
	if err != nil {
		return nil, fmt.Errorf("unable to resolve assertion key from key store for signing VP (did=%s): %w", *signerDID, err)
	}

	if validateVC {
		for _, cred := range credentials {
			err := h.verifier.VerifySignature(cred, &options.ProofOptions.Created)
			if err != nil {
				return nil, core.InvalidInputError("invalid credential (id=%s): %w", cred.ID, err)
			}
		}
	}

	switch options.Format {
	case JWTPresentationFormat:
		return h.buildJWTPresentation(ctx, *signerDID, credentials, options, key)
	case "":
		fallthrough
	case JSONLDPresentationFormat:
		return h.buildJSONLDPresentation(ctx, *signerDID, credentials, options, key)
	default:
		return nil, fmt.Errorf("unsupported presentation proof format: %s", options.Format)
	}
}

// buildJWTPresentation builds a JWT presentation according to https://www.w3.org/TR/vc-data-model/#json-web-token
func (h wallet) buildJWTPresentation(ctx context.Context, subjectDID did.DID, credentials []vc.VerifiableCredential, options PresentationOptions, key crypto.Key) (*vc.VerifiablePresentation, error) {
	headers := map[string]interface{}{
		jws.TypeKey: "JWT",
	}
	id := did.DIDURL{DID: subjectDID}
	id.Fragment = strings.ToLower(uuid.NewString())
	claims := map[string]interface{}{
		jwt.SubjectKey: subjectDID.String(),
		jwt.JwtIDKey:   id.String(),
		"vp": vc.VerifiablePresentation{
			Context:              append([]ssi.URI{VerifiableCredentialLDContextV1}, options.AdditionalContexts...),
			Type:                 append([]ssi.URI{VerifiablePresentationLDType}, options.AdditionalTypes...),
			VerifiableCredential: credentials,
		},
	}
	if options.ProofOptions.Nonce != nil {
		claims["nonce"] = *options.ProofOptions.Nonce
	}
	if options.ProofOptions.Domain != nil {
		claims[jwt.AudienceKey] = *options.ProofOptions.Domain
	}
	if options.ProofOptions.Created.IsZero() {
		claims[jwt.NotBeforeKey] = time.Now().Unix()
	} else {
		claims[jwt.NotBeforeKey] = int(options.ProofOptions.Created.Unix())
	}
	if options.ProofOptions.Expires != nil {
		claims[jwt.ExpirationKey] = int(options.ProofOptions.Expires.Unix())
	}
	for claimName, value := range options.ProofOptions.AdditionalProperties {
		claims[claimName] = value
	}
	token, err := h.keyStore.SignJWT(ctx, claims, headers, key)
	if err != nil {
		return nil, fmt.Errorf("unable to sign JWT presentation: %w", err)
	}
	return vc.ParseVerifiablePresentation(token)
}

func (h wallet) buildJSONLDPresentation(ctx context.Context, subjectDID did.DID, credentials []vc.VerifiableCredential, options PresentationOptions, key crypto.Key) (*vc.VerifiablePresentation, error) {
	ldContext := []ssi.URI{VerifiableCredentialLDContextV1, signature.JSONWebSignature2020Context}
	ldContext = append(ldContext, options.AdditionalContexts...)
	types := []ssi.URI{VerifiablePresentationLDType}
	types = append(types, options.AdditionalTypes...)

	id := did.DIDURL{DID: subjectDID}
	id.Fragment = strings.ToLower(uuid.NewString())
	idURI := id.URI()
	unsignedVP := &vc.VerifiablePresentation{
		ID:                   &idURI,
		Context:              ldContext,
		Type:                 types,
		VerifiableCredential: credentials,
	}

	// Convert to map[string]interface{} for signing
	documentBytes, err := unsignedVP.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var document proof.Document
	err = json.Unmarshal(documentBytes, &document)
	if err != nil {
		return nil, err
	}

	ldProof := proof.NewLDProof(options.ProofOptions)
	signingResult, err := ldProof.
		Sign(ctx, document, signature.JSONWebSignature2020{ContextLoader: h.jsonldManager.DocumentLoader(), Signer: h.keyStore}, key)
	if err != nil {
		return nil, fmt.Errorf("unable to sign VP with LD proof: %w", err)
	}
	resultJSON, _ := json.Marshal(signingResult)
	return vc.ParseVerifiablePresentation(string(resultJSON))
}

func (h wallet) Put(ctx context.Context, credentials ...vc.VerifiableCredential) error {
	err := h.walletStore.Write(ctx, func(tx stoabs.WriteTx) error {
		stats := tx.GetShelfWriter(statsShelf)
		var newCredentials uint32
		for _, curr := range credentials {
			subjectDID, err := curr.SubjectDID()
			if err != nil {
				return fmt.Errorf("unable to resolve subject DID from VC %s: %w", curr.ID, err)
			}
			walletKey := stoabs.BytesKey(curr.ID.String())
			// First check if the VC doesn't already exist; otherwise stats will be incorrect
			walletShelf := tx.GetShelfWriter(subjectDID.String())
			_, err = walletShelf.Get(walletKey)
			if err == nil {
				// Already exists
				continue
			} else if !errors.Is(err, stoabs.ErrKeyNotFound) {
				// Other error
				return fmt.Errorf("unable to check if credential %s already exists: %w", curr.ID, err)
			}
			// Write credential
			data, _ := curr.MarshalJSON()
			err = walletShelf.Put(walletKey, data)
			if err != nil {
				return fmt.Errorf("unable to store credential %s: %w", curr.ID, err)
			}
			newCredentials++
		}
		// Update stats
		currentCount, err := h.readCredentialCount(stats)
		if err != nil {
			return fmt.Errorf("unable to read wallet credential count: %w", err)
		}
		return stats.Put(credentialCountStatsKey, binary.BigEndian.AppendUint32([]byte{}, currentCount+newCredentials))
	}, stoabs.WithWriteLock()) // lock required for stats consistency
	if err != nil {
		return fmt.Errorf("unable to store credential(s): %w", err)
	}
	return nil
}

func (h wallet) List(ctx context.Context, holderDID did.DID) ([]vc.VerifiableCredential, error) {
	var result []vc.VerifiableCredential
	err := h.walletStore.ReadShelf(ctx, holderDID.String(), func(reader stoabs.Reader) error {
		return reader.Iterate(func(key stoabs.Key, value []byte) error {
			var cred vc.VerifiableCredential
			err := json.Unmarshal(value, &cred)
			if err != nil {
				return fmt.Errorf("unable to unmarshal credential %s: %w", string(key.Bytes()), err)
			}
			result = append(result, cred)
			return nil
		}, stoabs.BytesKey{})
	})
	if err != nil {
		return nil, fmt.Errorf("unable to list credentials: %w", err)
	}
	return result, nil
}

func (h wallet) Diagnostics() []core.DiagnosticResult {
	ctx := context.Background()
	var count uint32
	var err error
	err = h.walletStore.Read(ctx, func(tx stoabs.ReadTx) error {
		count, err = h.readCredentialCount(tx.GetShelfReader(statsShelf))
		return err
	})
	if err != nil {
		log.Logger().WithError(err).Warn("unable to read credential count in wallet")
	}
	return []core.DiagnosticResult{
		core.GenericDiagnosticResult{
			Title:   "credential_count",
			Outcome: int(count),
		},
	}
}

func (h wallet) IsEmpty() (bool, error) {
	ctx := context.Background()
	var count uint32
	var err error
	err = h.walletStore.Read(ctx, func(tx stoabs.ReadTx) error {
		count, err = h.readCredentialCount(tx.GetShelfReader(statsShelf))
		return err
	})
	return count == 0, err
}

func (h wallet) readCredentialCount(statsShelf stoabs.Reader) (uint32, error) {
	countBytes, err := statsShelf.Get(credentialCountStatsKey)
	if errors.Is(err, stoabs.ErrKeyNotFound) {
		// No stats yet
		countBytes = make([]byte, 4)
	} else if err != nil {
		return 0, fmt.Errorf("error reading credential count for wallet: %w", err)
	}
	return binary.BigEndian.Uint32(countBytes), nil
}
