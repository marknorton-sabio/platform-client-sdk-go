package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Keyperformanceindicator
type Keyperformanceindicator struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Key Performance Indicator.
	Name *string `json:"name,omitempty"`


	// OptimizationType - The optimization type of the Key Performance Indicator.
	OptimizationType *string `json:"optimizationType,omitempty"`


	// DateCreated - DateTime indicating when the Key Performance Indicator was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - DateTime indicating when the Key Performance Indicator was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Description - The description of the Key Performance Indicator.
	Description *string `json:"description,omitempty"`


	// KpiType - The type of Key Performance Indicator.
	KpiType *string `json:"kpiType,omitempty"`


	// Source - Source of values for Key Performance Indicator.
	Source *string `json:"source,omitempty"`


	// WrapUpCodeConfig - Defines what wrap up codes are mapped to Key Performance Indicator.
	WrapUpCodeConfig *Wrapupcodeconfig `json:"wrapUpCodeConfig,omitempty"`


	// Status - The status of the Key Performance Indicator.
	Status *string `json:"status,omitempty"`


	// KpiGroup - The group the Key Performance Indicator belongs to.
	KpiGroup *string `json:"kpiGroup,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Keyperformanceindicator) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Keyperformanceindicator
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		OptimizationType *string `json:"optimizationType,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		KpiType *string `json:"kpiType,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		WrapUpCodeConfig *Wrapupcodeconfig `json:"wrapUpCodeConfig,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		KpiGroup *string `json:"kpiGroup,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		OptimizationType: o.OptimizationType,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Description: o.Description,
		
		KpiType: o.KpiType,
		
		Source: o.Source,
		
		WrapUpCodeConfig: o.WrapUpCodeConfig,
		
		Status: o.Status,
		
		KpiGroup: o.KpiGroup,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Keyperformanceindicator) UnmarshalJSON(b []byte) error {
	var KeyperformanceindicatorMap map[string]interface{}
	err := json.Unmarshal(b, &KeyperformanceindicatorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KeyperformanceindicatorMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KeyperformanceindicatorMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if OptimizationType, ok := KeyperformanceindicatorMap["optimizationType"].(string); ok {
		o.OptimizationType = &OptimizationType
	}
    
	if dateCreatedString, ok := KeyperformanceindicatorMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KeyperformanceindicatorMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Description, ok := KeyperformanceindicatorMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if KpiType, ok := KeyperformanceindicatorMap["kpiType"].(string); ok {
		o.KpiType = &KpiType
	}
    
	if Source, ok := KeyperformanceindicatorMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if WrapUpCodeConfig, ok := KeyperformanceindicatorMap["wrapUpCodeConfig"].(map[string]interface{}); ok {
		WrapUpCodeConfigString, _ := json.Marshal(WrapUpCodeConfig)
		json.Unmarshal(WrapUpCodeConfigString, &o.WrapUpCodeConfig)
	}
	
	if Status, ok := KeyperformanceindicatorMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if KpiGroup, ok := KeyperformanceindicatorMap["kpiGroup"].(string); ok {
		o.KpiGroup = &KpiGroup
	}
    
	if SelfUri, ok := KeyperformanceindicatorMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Keyperformanceindicator) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
