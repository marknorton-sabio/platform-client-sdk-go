package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Addworkplanrotationrequest
type Addworkplanrotationrequest struct { 
	// Name - Name of this work plan rotation
	Name *string `json:"name,omitempty"`


	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Agents - Agents in this work plan rotation
	Agents *[]Addworkplanrotationagentrequest `json:"agents,omitempty"`


	// Pattern - Pattern with list of work plan IDs that rotate on a weekly basis
	Pattern *Workplanpatternrequest `json:"pattern,omitempty"`

}

func (o *Addworkplanrotationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addworkplanrotationrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Agents *[]Addworkplanrotationagentrequest `json:"agents,omitempty"`
		
		Pattern *Workplanpatternrequest `json:"pattern,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		DateRange: o.DateRange,
		
		Agents: o.Agents,
		
		Pattern: o.Pattern,
		Alias:    (*Alias)(o),
	})
}

func (o *Addworkplanrotationrequest) UnmarshalJSON(b []byte) error {
	var AddworkplanrotationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddworkplanrotationrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AddworkplanrotationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if DateRange, ok := AddworkplanrotationrequestMap["dateRange"].(map[string]interface{}); ok {
		DateRangeString, _ := json.Marshal(DateRange)
		json.Unmarshal(DateRangeString, &o.DateRange)
	}
	
	if Agents, ok := AddworkplanrotationrequestMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if Pattern, ok := AddworkplanrotationrequestMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Addworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
