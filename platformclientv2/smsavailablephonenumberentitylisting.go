package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsavailablephonenumberentitylisting
type Smsavailablephonenumberentitylisting struct { 
	// Entities
	Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`

}

func (o *Smsavailablephonenumberentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsavailablephonenumberentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsavailablephonenumberentitylisting) UnmarshalJSON(b []byte) error {
	var SmsavailablephonenumberentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &SmsavailablephonenumberentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := SmsavailablephonenumberentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumberentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
