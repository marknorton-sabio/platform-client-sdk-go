package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Customerinteractioncenter
type Customerinteractioncenter struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// IssuerURI
	IssuerURI *string `json:"issuerURI,omitempty"`


	// SsoTargetURI
	SsoTargetURI *string `json:"ssoTargetURI,omitempty"`


	// SloURI
	SloURI *string `json:"sloURI,omitempty"`


	// SloBinding
	SloBinding *string `json:"sloBinding,omitempty"`


	// RelyingPartyIdentifier
	RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`


	// Certificate
	Certificate *string `json:"certificate,omitempty"`


	// Certificates
	Certificates *[]string `json:"certificates,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Customerinteractioncenter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Customerinteractioncenter
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		IssuerURI *string `json:"issuerURI,omitempty"`
		
		SsoTargetURI *string `json:"ssoTargetURI,omitempty"`
		
		SloURI *string `json:"sloURI,omitempty"`
		
		SloBinding *string `json:"sloBinding,omitempty"`
		
		RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`
		
		Certificate *string `json:"certificate,omitempty"`
		
		Certificates *[]string `json:"certificates,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Disabled: o.Disabled,
		
		IssuerURI: o.IssuerURI,
		
		SsoTargetURI: o.SsoTargetURI,
		
		SloURI: o.SloURI,
		
		SloBinding: o.SloBinding,
		
		RelyingPartyIdentifier: o.RelyingPartyIdentifier,
		
		Certificate: o.Certificate,
		
		Certificates: o.Certificates,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Customerinteractioncenter) UnmarshalJSON(b []byte) error {
	var CustomerinteractioncenterMap map[string]interface{}
	err := json.Unmarshal(b, &CustomerinteractioncenterMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CustomerinteractioncenterMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CustomerinteractioncenterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Disabled, ok := CustomerinteractioncenterMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if IssuerURI, ok := CustomerinteractioncenterMap["issuerURI"].(string); ok {
		o.IssuerURI = &IssuerURI
	}
    
	if SsoTargetURI, ok := CustomerinteractioncenterMap["ssoTargetURI"].(string); ok {
		o.SsoTargetURI = &SsoTargetURI
	}
    
	if SloURI, ok := CustomerinteractioncenterMap["sloURI"].(string); ok {
		o.SloURI = &SloURI
	}
    
	if SloBinding, ok := CustomerinteractioncenterMap["sloBinding"].(string); ok {
		o.SloBinding = &SloBinding
	}
    
	if RelyingPartyIdentifier, ok := CustomerinteractioncenterMap["relyingPartyIdentifier"].(string); ok {
		o.RelyingPartyIdentifier = &RelyingPartyIdentifier
	}
    
	if Certificate, ok := CustomerinteractioncenterMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
    
	if Certificates, ok := CustomerinteractioncenterMap["certificates"].([]interface{}); ok {
		CertificatesString, _ := json.Marshal(Certificates)
		json.Unmarshal(CertificatesString, &o.Certificates)
	}
	
	if SelfUri, ok := CustomerinteractioncenterMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Customerinteractioncenter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
