package platformclientv2
import (
	"encoding/json"
)

// Credential
type Credential struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType - The type of credential.
	VarType *Credentialtype `json:"type,omitempty"`


	// CredentialFields
	CredentialFields *map[string]string `json:"credentialFields,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Credential) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
