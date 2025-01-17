package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiccontact
type Externalcontactsunresolvedcontactchangedtopiccontact struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// FirstName
	FirstName *string `json:"firstName,omitempty"`


	// MiddleName
	MiddleName *string `json:"middleName,omitempty"`


	// LastName
	LastName *string `json:"lastName,omitempty"`


	// Salutation
	Salutation *string `json:"salutation,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// WorkPhone
	WorkPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"workPhone,omitempty"`


	// CellPhone
	CellPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"cellPhone,omitempty"`


	// HomePhone
	HomePhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"homePhone,omitempty"`


	// OtherPhone
	OtherPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"otherPhone,omitempty"`


	// WorkEmail
	WorkEmail *string `json:"workEmail,omitempty"`


	// PersonalEmail
	PersonalEmail *string `json:"personalEmail,omitempty"`


	// OtherEmail
	OtherEmail *string `json:"otherEmail,omitempty"`


	// Address
	Address *Externalcontactsunresolvedcontactchangedtopiccontactaddress `json:"address,omitempty"`


	// SurveyOptOut
	SurveyOptOut *bool `json:"surveyOptOut,omitempty"`


	// ExternalSystemUrl
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`


	// TwitterId
	TwitterId *Externalcontactsunresolvedcontactchangedtopictwitterid `json:"twitterId,omitempty"`


	// LineId
	LineId *Externalcontactsunresolvedcontactchangedtopiclineid `json:"lineId,omitempty"`


	// WhatsAppId
	WhatsAppId *Externalcontactsunresolvedcontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`


	// FacebookId
	FacebookId *Externalcontactsunresolvedcontactchangedtopicfacebookid `json:"facebookId,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopiccontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopiccontact
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		MiddleName *string `json:"middleName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Salutation *string `json:"salutation,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		WorkPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"workPhone,omitempty"`
		
		CellPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"cellPhone,omitempty"`
		
		HomePhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"homePhone,omitempty"`
		
		OtherPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"otherPhone,omitempty"`
		
		WorkEmail *string `json:"workEmail,omitempty"`
		
		PersonalEmail *string `json:"personalEmail,omitempty"`
		
		OtherEmail *string `json:"otherEmail,omitempty"`
		
		Address *Externalcontactsunresolvedcontactchangedtopiccontactaddress `json:"address,omitempty"`
		
		SurveyOptOut *bool `json:"surveyOptOut,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		TwitterId *Externalcontactsunresolvedcontactchangedtopictwitterid `json:"twitterId,omitempty"`
		
		LineId *Externalcontactsunresolvedcontactchangedtopiclineid `json:"lineId,omitempty"`
		
		WhatsAppId *Externalcontactsunresolvedcontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`
		
		FacebookId *Externalcontactsunresolvedcontactchangedtopicfacebookid `json:"facebookId,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		FirstName: o.FirstName,
		
		MiddleName: o.MiddleName,
		
		LastName: o.LastName,
		
		Salutation: o.Salutation,
		
		Title: o.Title,
		
		WorkPhone: o.WorkPhone,
		
		CellPhone: o.CellPhone,
		
		HomePhone: o.HomePhone,
		
		OtherPhone: o.OtherPhone,
		
		WorkEmail: o.WorkEmail,
		
		PersonalEmail: o.PersonalEmail,
		
		OtherEmail: o.OtherEmail,
		
		Address: o.Address,
		
		SurveyOptOut: o.SurveyOptOut,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		TwitterId: o.TwitterId,
		
		LineId: o.LineId,
		
		WhatsAppId: o.WhatsAppId,
		
		FacebookId: o.FacebookId,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiccontact) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiccontactMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiccontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if FirstName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if MiddleName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["middleName"].(string); ok {
		o.MiddleName = &MiddleName
	}
    
	if LastName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Salutation, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["salutation"].(string); ok {
		o.Salutation = &Salutation
	}
    
	if Title, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if WorkPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["workPhone"].(map[string]interface{}); ok {
		WorkPhoneString, _ := json.Marshal(WorkPhone)
		json.Unmarshal(WorkPhoneString, &o.WorkPhone)
	}
	
	if CellPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["cellPhone"].(map[string]interface{}); ok {
		CellPhoneString, _ := json.Marshal(CellPhone)
		json.Unmarshal(CellPhoneString, &o.CellPhone)
	}
	
	if HomePhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["homePhone"].(map[string]interface{}); ok {
		HomePhoneString, _ := json.Marshal(HomePhone)
		json.Unmarshal(HomePhoneString, &o.HomePhone)
	}
	
	if OtherPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["otherPhone"].(map[string]interface{}); ok {
		OtherPhoneString, _ := json.Marshal(OtherPhone)
		json.Unmarshal(OtherPhoneString, &o.OtherPhone)
	}
	
	if WorkEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["workEmail"].(string); ok {
		o.WorkEmail = &WorkEmail
	}
    
	if PersonalEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["personalEmail"].(string); ok {
		o.PersonalEmail = &PersonalEmail
	}
    
	if OtherEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["otherEmail"].(string); ok {
		o.OtherEmail = &OtherEmail
	}
    
	if Address, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if SurveyOptOut, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["surveyOptOut"].(bool); ok {
		o.SurveyOptOut = &SurveyOptOut
	}
    
	if ExternalSystemUrl, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if TwitterId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if LineId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["lineId"].(map[string]interface{}); ok {
		LineIdString, _ := json.Marshal(LineId)
		json.Unmarshal(LineIdString, &o.LineId)
	}
	
	if WhatsAppId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["whatsAppId"].(map[string]interface{}); ok {
		WhatsAppIdString, _ := json.Marshal(WhatsAppId)
		json.Unmarshal(WhatsAppIdString, &o.WhatsAppId)
	}
	
	if FacebookId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["facebookId"].(map[string]interface{}); ok {
		FacebookIdString, _ := json.Marshal(FacebookId)
		json.Unmarshal(FacebookIdString, &o.FacebookId)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiccontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
