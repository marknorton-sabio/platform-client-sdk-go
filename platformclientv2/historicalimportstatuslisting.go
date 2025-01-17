package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalimportstatuslisting
type Historicalimportstatuslisting struct { 
	// Entities
	Entities *[]Historicalimportstatus `json:"entities,omitempty"`

}

func (o *Historicalimportstatuslisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalimportstatuslisting
	
	return json.Marshal(&struct { 
		Entities *[]Historicalimportstatus `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalimportstatuslisting) UnmarshalJSON(b []byte) error {
	var HistoricalimportstatuslistingMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalimportstatuslistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := HistoricalimportstatuslistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalimportstatuslisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
