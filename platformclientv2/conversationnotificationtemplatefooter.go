package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnotificationtemplatefooter - Template footer object.
type Conversationnotificationtemplatefooter struct { 
	// Text - Footer text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`

}

func (o *Conversationnotificationtemplatefooter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationnotificationtemplatefooter
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationnotificationtemplatefooter) UnmarshalJSON(b []byte) error {
	var ConversationnotificationtemplatefooterMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationnotificationtemplatefooterMap)
	if err != nil {
		return err
	}
	
	if Text, ok := ConversationnotificationtemplatefooterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatefooter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
