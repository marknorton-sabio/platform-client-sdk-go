package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicdisconnectreason
type Conversationeventtopicdisconnectreason struct { 
	// VarType - Disconnect reason protocol type.
	VarType *string `json:"type,omitempty"`


	// Code - Protocol specific reason code. See the Q.850 and SIP specs.
	Code *int `json:"code,omitempty"`


	// Phrase - Human readable English description of the disconnect reason.
	Phrase *string `json:"phrase,omitempty"`

}

func (o *Conversationeventtopicdisconnectreason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicdisconnectreason
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Code *int `json:"code,omitempty"`
		
		Phrase *string `json:"phrase,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Code: o.Code,
		
		Phrase: o.Phrase,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicdisconnectreason) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicdisconnectreasonMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicdisconnectreasonMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationeventtopicdisconnectreasonMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Code, ok := ConversationeventtopicdisconnectreasonMap["code"].(float64); ok {
		CodeInt := int(Code)
		o.Code = &CodeInt
	}
	
	if Phrase, ok := ConversationeventtopicdisconnectreasonMap["phrase"].(string); ok {
		o.Phrase = &Phrase
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
