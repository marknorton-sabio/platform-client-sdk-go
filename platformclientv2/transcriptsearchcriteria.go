package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptsearchcriteria
type Transcriptsearchcriteria struct { 
	// EndValue - The end value of the range. This field is used for range search types.
	EndValue *string `json:"endValue,omitempty"`


	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`


	// StartValue - The start value of the range. This field is used for range search types.
	StartValue *string `json:"startValue,omitempty"`


	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`


	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`


	// Group - Groups multiple conditions
	Group *[]Transcriptsearchcriteria `json:"group,omitempty"`


	// DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat *string `json:"dateFormat,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Fields - Field names to search against
	Fields *[]string `json:"fields,omitempty"`

}

func (o *Transcriptsearchcriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptsearchcriteria
	
	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Group *[]Transcriptsearchcriteria `json:"group,omitempty"`
		
		DateFormat *string `json:"dateFormat,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		*Alias
	}{ 
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Value: o.Value,
		
		Operator: o.Operator,
		
		Group: o.Group,
		
		DateFormat: o.DateFormat,
		
		VarType: o.VarType,
		
		Fields: o.Fields,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptsearchcriteria) UnmarshalJSON(b []byte) error {
	var TranscriptsearchcriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptsearchcriteriaMap)
	if err != nil {
		return err
	}
	
	if EndValue, ok := TranscriptsearchcriteriaMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
    
	if Values, ok := TranscriptsearchcriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := TranscriptsearchcriteriaMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
    
	if Value, ok := TranscriptsearchcriteriaMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Operator, ok := TranscriptsearchcriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Group, ok := TranscriptsearchcriteriaMap["group"].([]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if DateFormat, ok := TranscriptsearchcriteriaMap["dateFormat"].(string); ok {
		o.DateFormat = &DateFormat
	}
    
	if VarType, ok := TranscriptsearchcriteriaMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Fields, ok := TranscriptsearchcriteriaMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptsearchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
