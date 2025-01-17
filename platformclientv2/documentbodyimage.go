package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyimage
type Documentbodyimage struct { 
	// Url - The URL for the image.
	Url *string `json:"url,omitempty"`

}

func (o *Documentbodyimage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodyimage
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodyimage) UnmarshalJSON(b []byte) error {
	var DocumentbodyimageMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyimageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := DocumentbodyimageMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyimage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
