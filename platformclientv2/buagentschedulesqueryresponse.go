package platformclientv2
import (
	"encoding/json"
)

// Buagentschedulesqueryresponse
type Buagentschedulesqueryresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulequeryresponse `json:"agentSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulesqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}