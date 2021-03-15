package credential

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNutsOrganizationCredentialSubject(t *testing.T) {
	subject := NutsOrganizationCredentialSubject{
		ID: "123456",
		Organization: map[string]string{
			"name": "Awesome B.V.",
			"city": "Somewhere",
		},
	}
	actual, _ := json.Marshal(subject)
	assert.JSONEq(t, `{"id":"123456","organization":{"city":"Somewhere","name":"Awesome B.V."}}`, string(actual))
}
