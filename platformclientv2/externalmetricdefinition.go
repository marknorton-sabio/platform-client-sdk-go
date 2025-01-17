package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdefinition
type Externalmetricdefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the External Metric Definition
	Name *string `json:"name,omitempty"`


	// Unit - The unit of the External Metric Definition
	Unit *string `json:"unit,omitempty"`


	// UnitDefinition - The unit definition of the External Metric Definition
	UnitDefinition *string `json:"unitDefinition,omitempty"`


	// Precision - The decimal precision of the External Metric Definition
	Precision *int `json:"precision,omitempty"`


	// DefaultObjectiveType - The default objective type of the External Metric Definition
	DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`


	// RetentionMonths - The retention in months of the External Metric Definition
	RetentionMonths *int `json:"retentionMonths,omitempty"`


	// Enabled - True if the External Metric Definition is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// InUse - True if the External Metric Definition is in use
	InUse *bool `json:"inUse,omitempty"`


	// DateLastRefreshed - The last date and time that the metric data was refreshed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastRefreshed *time.Time `json:"dateLastRefreshed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Externalmetricdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdefinition
	
	DateLastRefreshed := new(string)
	if o.DateLastRefreshed != nil {
		
		*DateLastRefreshed = timeutil.Strftime(o.DateLastRefreshed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastRefreshed = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Unit *string `json:"unit,omitempty"`
		
		UnitDefinition *string `json:"unitDefinition,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`
		
		RetentionMonths *int `json:"retentionMonths,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InUse *bool `json:"inUse,omitempty"`
		
		DateLastRefreshed *string `json:"dateLastRefreshed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Unit: o.Unit,
		
		UnitDefinition: o.UnitDefinition,
		
		Precision: o.Precision,
		
		DefaultObjectiveType: o.DefaultObjectiveType,
		
		RetentionMonths: o.RetentionMonths,
		
		Enabled: o.Enabled,
		
		InUse: o.InUse,
		
		DateLastRefreshed: DateLastRefreshed,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdefinition) UnmarshalJSON(b []byte) error {
	var ExternalmetricdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalmetricdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ExternalmetricdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Unit, ok := ExternalmetricdefinitionMap["unit"].(string); ok {
		o.Unit = &Unit
	}
    
	if UnitDefinition, ok := ExternalmetricdefinitionMap["unitDefinition"].(string); ok {
		o.UnitDefinition = &UnitDefinition
	}
    
	if Precision, ok := ExternalmetricdefinitionMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if DefaultObjectiveType, ok := ExternalmetricdefinitionMap["defaultObjectiveType"].(string); ok {
		o.DefaultObjectiveType = &DefaultObjectiveType
	}
    
	if RetentionMonths, ok := ExternalmetricdefinitionMap["retentionMonths"].(float64); ok {
		RetentionMonthsInt := int(RetentionMonths)
		o.RetentionMonths = &RetentionMonthsInt
	}
	
	if Enabled, ok := ExternalmetricdefinitionMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InUse, ok := ExternalmetricdefinitionMap["inUse"].(bool); ok {
		o.InUse = &InUse
	}
    
	if dateLastRefreshedString, ok := ExternalmetricdefinitionMap["dateLastRefreshed"].(string); ok {
		DateLastRefreshed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastRefreshedString)
		o.DateLastRefreshed = &DateLastRefreshed
	}
	
	if SelfUri, ok := ExternalmetricdefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
