package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemseventsnotificationcustomattribute
type Workitemseventsnotificationcustomattribute struct { 
	// DataType
	DataType *string `json:"dataType,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *Workitemseventsnotificationcustomattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workitemseventsnotificationcustomattribute
	
	return json.Marshal(&struct { 
		DataType *string `json:"dataType,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		DataType: o.DataType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Workitemseventsnotificationcustomattribute) UnmarshalJSON(b []byte) error {
	var WorkitemseventsnotificationcustomattributeMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemseventsnotificationcustomattributeMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := WorkitemseventsnotificationcustomattributeMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    
	if Value, ok := WorkitemseventsnotificationcustomattributeMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemseventsnotificationcustomattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
