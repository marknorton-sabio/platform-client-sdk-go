package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentpositionproperties
type Contentpositionproperties struct { 
	// Top - Top positioning offset.
	Top *string `json:"top,omitempty"`


	// Bottom - Bottom positioning offset.
	Bottom *string `json:"bottom,omitempty"`


	// Left - Left positioning offset.
	Left *string `json:"left,omitempty"`


	// Right - Right positioning offset.
	Right *string `json:"right,omitempty"`

}

func (o *Contentpositionproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentpositionproperties
	
	return json.Marshal(&struct { 
		Top *string `json:"top,omitempty"`
		
		Bottom *string `json:"bottom,omitempty"`
		
		Left *string `json:"left,omitempty"`
		
		Right *string `json:"right,omitempty"`
		*Alias
	}{ 
		Top: o.Top,
		
		Bottom: o.Bottom,
		
		Left: o.Left,
		
		Right: o.Right,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentpositionproperties) UnmarshalJSON(b []byte) error {
	var ContentpositionpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ContentpositionpropertiesMap)
	if err != nil {
		return err
	}
	
	if Top, ok := ContentpositionpropertiesMap["top"].(string); ok {
		o.Top = &Top
	}
    
	if Bottom, ok := ContentpositionpropertiesMap["bottom"].(string); ok {
		o.Bottom = &Bottom
	}
    
	if Left, ok := ContentpositionpropertiesMap["left"].(string); ok {
		o.Left = &Left
	}
    
	if Right, ok := ContentpositionpropertiesMap["right"].(string); ok {
		o.Right = &Right
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentpositionproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
