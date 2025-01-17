package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineid - User information for a Line account
type Lineid struct { 
	// Ids - The set of Line userIds that this person has. Each userId is specific to the Line channel that the user interacts with.
	Ids *[]Lineuserid `json:"ids,omitempty"`


	// DisplayName - The displayName of this person's account in Line
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Lineid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineid
	
	return json.Marshal(&struct { 
		Ids *[]Lineuserid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Lineid) UnmarshalJSON(b []byte) error {
	var LineidMap map[string]interface{}
	err := json.Unmarshal(b, &LineidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := LineidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := LineidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lineid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
