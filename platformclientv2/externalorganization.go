package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganization
type Externalorganization struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the company.
	Name *string `json:"name,omitempty"`


	// CompanyType
	CompanyType *string `json:"companyType,omitempty"`


	// Industry
	Industry *string `json:"industry,omitempty"`


	// PrimaryContactId
	PrimaryContactId *string `json:"primaryContactId,omitempty"`


	// Address
	Address *Contactaddress `json:"address,omitempty"`


	// PhoneNumber
	PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`


	// FaxNumber
	FaxNumber *Phonenumber `json:"faxNumber,omitempty"`


	// EmployeeCount
	EmployeeCount *int `json:"employeeCount,omitempty"`


	// Revenue
	Revenue *int `json:"revenue,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// Websites
	Websites *[]string `json:"websites,omitempty"`


	// Tickers
	Tickers *[]Ticker `json:"tickers,omitempty"`


	// TwitterId
	TwitterId *Twitterid `json:"twitterId,omitempty"`


	// ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the organization. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`


	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`


	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`


	// Trustor
	Trustor *Trustor `json:"trustor,omitempty"`


	// Schema - The schema defining custom fields for this contact
	Schema *Dataschema `json:"schema,omitempty"`


	// CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`


	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Externalorganization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalorganization
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CompanyType *string `json:"companyType,omitempty"`
		
		Industry *string `json:"industry,omitempty"`
		
		PrimaryContactId *string `json:"primaryContactId,omitempty"`
		
		Address *Contactaddress `json:"address,omitempty"`
		
		PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`
		
		FaxNumber *Phonenumber `json:"faxNumber,omitempty"`
		
		EmployeeCount *int `json:"employeeCount,omitempty"`
		
		Revenue *int `json:"revenue,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Websites *[]string `json:"websites,omitempty"`
		
		Tickers *[]Ticker `json:"tickers,omitempty"`
		
		TwitterId *Twitterid `json:"twitterId,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		Trustor *Trustor `json:"trustor,omitempty"`
		
		Schema *Dataschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CompanyType: o.CompanyType,
		
		Industry: o.Industry,
		
		PrimaryContactId: o.PrimaryContactId,
		
		Address: o.Address,
		
		PhoneNumber: o.PhoneNumber,
		
		FaxNumber: o.FaxNumber,
		
		EmployeeCount: o.EmployeeCount,
		
		Revenue: o.Revenue,
		
		Tags: o.Tags,
		
		Websites: o.Websites,
		
		Tickers: o.Tickers,
		
		TwitterId: o.TwitterId,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		Trustor: o.Trustor,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		ExternalDataSources: o.ExternalDataSources,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalorganization) UnmarshalJSON(b []byte) error {
	var ExternalorganizationMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalorganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalorganizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ExternalorganizationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CompanyType, ok := ExternalorganizationMap["companyType"].(string); ok {
		o.CompanyType = &CompanyType
	}
    
	if Industry, ok := ExternalorganizationMap["industry"].(string); ok {
		o.Industry = &Industry
	}
    
	if PrimaryContactId, ok := ExternalorganizationMap["primaryContactId"].(string); ok {
		o.PrimaryContactId = &PrimaryContactId
	}
    
	if Address, ok := ExternalorganizationMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if PhoneNumber, ok := ExternalorganizationMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if FaxNumber, ok := ExternalorganizationMap["faxNumber"].(map[string]interface{}); ok {
		FaxNumberString, _ := json.Marshal(FaxNumber)
		json.Unmarshal(FaxNumberString, &o.FaxNumber)
	}
	
	if EmployeeCount, ok := ExternalorganizationMap["employeeCount"].(float64); ok {
		EmployeeCountInt := int(EmployeeCount)
		o.EmployeeCount = &EmployeeCountInt
	}
	
	if Revenue, ok := ExternalorganizationMap["revenue"].(float64); ok {
		RevenueInt := int(Revenue)
		o.Revenue = &RevenueInt
	}
	
	if Tags, ok := ExternalorganizationMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Websites, ok := ExternalorganizationMap["websites"].([]interface{}); ok {
		WebsitesString, _ := json.Marshal(Websites)
		json.Unmarshal(WebsitesString, &o.Websites)
	}
	
	if Tickers, ok := ExternalorganizationMap["tickers"].([]interface{}); ok {
		TickersString, _ := json.Marshal(Tickers)
		json.Unmarshal(TickersString, &o.Tickers)
	}
	
	if TwitterId, ok := ExternalorganizationMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if ExternalSystemUrl, ok := ExternalorganizationMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if modifyDateString, ok := ExternalorganizationMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	
	if createDateString, ok := ExternalorganizationMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if Trustor, ok := ExternalorganizationMap["trustor"].(map[string]interface{}); ok {
		TrustorString, _ := json.Marshal(Trustor)
		json.Unmarshal(TrustorString, &o.Trustor)
	}
	
	if Schema, ok := ExternalorganizationMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalorganizationMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if ExternalDataSources, ok := ExternalorganizationMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if SelfUri, ok := ExternalorganizationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
