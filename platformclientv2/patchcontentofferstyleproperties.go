package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcontentofferstyleproperties
type Patchcontentofferstyleproperties struct { 
	// Padding - Padding of the offer. (eg. 10px)
	Padding *string `json:"padding,omitempty"`


	// Color - Text color of the offer. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// BackgroundColor - Background color of the offer. (eg. #000000)
	BackgroundColor *string `json:"backgroundColor,omitempty"`

}

func (o *Patchcontentofferstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchcontentofferstyleproperties
	
	return json.Marshal(&struct { 
		Padding *string `json:"padding,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		*Alias
	}{ 
		Padding: o.Padding,
		
		Color: o.Color,
		
		BackgroundColor: o.BackgroundColor,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchcontentofferstyleproperties) UnmarshalJSON(b []byte) error {
	var PatchcontentofferstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &PatchcontentofferstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Padding, ok := PatchcontentofferstylepropertiesMap["padding"].(string); ok {
		o.Padding = &Padding
	}
    
	if Color, ok := PatchcontentofferstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if BackgroundColor, ok := PatchcontentofferstylepropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchcontentofferstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
