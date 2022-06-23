package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Extendedwrapup
type Extendedwrapup struct { 
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`


	// Name - The user configured wrap up code name.
	Name *string `json:"name,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The length of time in seconds that the agent spent doing after call work.
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrapup was finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// Provisional - Indicates if this is a pending save and should not require a code to be specified.  This allows someone to save some temporary wrapup that will be used later.
	Provisional *bool `json:"provisional,omitempty"`


	// DisableEndTimeUpdates - Prevent updates to wrapup end time when set to true.
	DisableEndTimeUpdates *bool `json:"disableEndTimeUpdates,omitempty"`

}

func (o *Extendedwrapup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Extendedwrapup
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Provisional *bool `json:"provisional,omitempty"`
		
		DisableEndTimeUpdates *bool `json:"disableEndTimeUpdates,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Name: o.Name,
		
		Notes: o.Notes,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		EndTime: EndTime,
		
		Provisional: o.Provisional,
		
		DisableEndTimeUpdates: o.DisableEndTimeUpdates,
		Alias:    (*Alias)(o),
	})
}

func (o *Extendedwrapup) UnmarshalJSON(b []byte) error {
	var ExtendedwrapupMap map[string]interface{}
	err := json.Unmarshal(b, &ExtendedwrapupMap)
	if err != nil {
		return err
	}
	
	if Code, ok := ExtendedwrapupMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Name, ok := ExtendedwrapupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Notes, ok := ExtendedwrapupMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if Tags, ok := ExtendedwrapupMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := ExtendedwrapupMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if endTimeString, ok := ExtendedwrapupMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if Provisional, ok := ExtendedwrapupMap["provisional"].(bool); ok {
		o.Provisional = &Provisional
	}
    
	if DisableEndTimeUpdates, ok := ExtendedwrapupMap["disableEndTimeUpdates"].(bool); ok {
		o.DisableEndTimeUpdates = &DisableEndTimeUpdates
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Extendedwrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}