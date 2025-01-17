package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagemediapolicyconditions
type Messagemediapolicyconditions struct { 
	// ForUsers
	ForUsers *[]User `json:"forUsers,omitempty"`


	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`


	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`


	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`


	// Languages
	Languages *[]Language `json:"languages,omitempty"`


	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`


	// CustomerParticipation
	CustomerParticipation *string `json:"customerParticipation,omitempty"`

}

func (o *Messagemediapolicyconditions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagemediapolicyconditions
	
	return json.Marshal(&struct { 
		ForUsers *[]User `json:"forUsers,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		Languages *[]Language `json:"languages,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		
		CustomerParticipation *string `json:"customerParticipation,omitempty"`
		*Alias
	}{ 
		ForUsers: o.ForUsers,
		
		DateRanges: o.DateRanges,
		
		ForQueues: o.ForQueues,
		
		WrapupCodes: o.WrapupCodes,
		
		Languages: o.Languages,
		
		TimeAllowed: o.TimeAllowed,
		
		CustomerParticipation: o.CustomerParticipation,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagemediapolicyconditions) UnmarshalJSON(b []byte) error {
	var MessagemediapolicyconditionsMap map[string]interface{}
	err := json.Unmarshal(b, &MessagemediapolicyconditionsMap)
	if err != nil {
		return err
	}
	
	if ForUsers, ok := MessagemediapolicyconditionsMap["forUsers"].([]interface{}); ok {
		ForUsersString, _ := json.Marshal(ForUsers)
		json.Unmarshal(ForUsersString, &o.ForUsers)
	}
	
	if DateRanges, ok := MessagemediapolicyconditionsMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	
	if ForQueues, ok := MessagemediapolicyconditionsMap["forQueues"].([]interface{}); ok {
		ForQueuesString, _ := json.Marshal(ForQueues)
		json.Unmarshal(ForQueuesString, &o.ForQueues)
	}
	
	if WrapupCodes, ok := MessagemediapolicyconditionsMap["wrapupCodes"].([]interface{}); ok {
		WrapupCodesString, _ := json.Marshal(WrapupCodes)
		json.Unmarshal(WrapupCodesString, &o.WrapupCodes)
	}
	
	if Languages, ok := MessagemediapolicyconditionsMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if TimeAllowed, ok := MessagemediapolicyconditionsMap["timeAllowed"].(map[string]interface{}); ok {
		TimeAllowedString, _ := json.Marshal(TimeAllowed)
		json.Unmarshal(TimeAllowedString, &o.TimeAllowed)
	}
	
	if CustomerParticipation, ok := MessagemediapolicyconditionsMap["customerParticipation"].(string); ok {
		o.CustomerParticipation = &CustomerParticipation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagemediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
