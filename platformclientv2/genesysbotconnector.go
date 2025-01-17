package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Genesysbotconnector
type Genesysbotconnector struct { 
	// QueryParameters - User defined name/value parameters passed to the BotConnector bot.
	QueryParameters *map[string]string `json:"queryParameters,omitempty"`

}

func (o *Genesysbotconnector) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Genesysbotconnector
	
	return json.Marshal(&struct { 
		QueryParameters *map[string]string `json:"queryParameters,omitempty"`
		*Alias
	}{ 
		QueryParameters: o.QueryParameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Genesysbotconnector) UnmarshalJSON(b []byte) error {
	var GenesysbotconnectorMap map[string]interface{}
	err := json.Unmarshal(b, &GenesysbotconnectorMap)
	if err != nil {
		return err
	}
	
	if QueryParameters, ok := GenesysbotconnectorMap["queryParameters"].(map[string]interface{}); ok {
		QueryParametersString, _ := json.Marshal(QueryParameters)
		json.Unmarshal(QueryParametersString, &o.QueryParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Genesysbotconnector) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
