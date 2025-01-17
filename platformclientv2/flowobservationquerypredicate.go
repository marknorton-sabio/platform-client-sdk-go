package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowobservationquerypredicate
type Flowobservationquerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for dimension predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for dimension predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

func (o *Flowobservationquerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowobservationquerypredicate
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowobservationquerypredicate) UnmarshalJSON(b []byte) error {
	var FlowobservationquerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &FlowobservationquerypredicateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := FlowobservationquerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := FlowobservationquerypredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Operator, ok := FlowobservationquerypredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := FlowobservationquerypredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := FlowobservationquerypredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowobservationquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
