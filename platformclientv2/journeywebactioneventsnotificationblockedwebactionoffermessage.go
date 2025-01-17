package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationblockedwebactionoffermessage
type Journeywebactioneventsnotificationblockedwebactionoffermessage struct { 
	// Action
	Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`


	// ActionMap
	ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`


	// ActionTarget
	ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`


	// BlockingReason
	BlockingReason *string `json:"blockingReason,omitempty"`


	// BlockingActionMap
	BlockingActionMap *Journeywebactioneventsnotificationactionmap `json:"blockingActionMap,omitempty"`


	// BlockingAction
	BlockingAction *Journeywebactioneventsnotificationeventaction `json:"blockingAction,omitempty"`


	// BlockingFrequencyCapBehaviour
	BlockingFrequencyCapBehaviour *string `json:"blockingFrequencyCapBehaviour,omitempty"`


	// BlockingPageUrlConditions
	BlockingPageUrlConditions *[]Journeywebactioneventsnotificationactionmappageurlcondition `json:"blockingPageUrlConditions,omitempty"`


	// BlockingScheduleGroup
	BlockingScheduleGroup *Journeywebactioneventsnotificationschedulegroup `json:"blockingScheduleGroup,omitempty"`


	// BlockingEmergencyScheduleGroup
	BlockingEmergencyScheduleGroup *Journeywebactioneventsnotificationemergencygroup `json:"blockingEmergencyScheduleGroup,omitempty"`

}

func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationblockedwebactionoffermessage
	
	return json.Marshal(&struct { 
		Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`
		
		ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		BlockingReason *string `json:"blockingReason,omitempty"`
		
		BlockingActionMap *Journeywebactioneventsnotificationactionmap `json:"blockingActionMap,omitempty"`
		
		BlockingAction *Journeywebactioneventsnotificationeventaction `json:"blockingAction,omitempty"`
		
		BlockingFrequencyCapBehaviour *string `json:"blockingFrequencyCapBehaviour,omitempty"`
		
		BlockingPageUrlConditions *[]Journeywebactioneventsnotificationactionmappageurlcondition `json:"blockingPageUrlConditions,omitempty"`
		
		BlockingScheduleGroup *Journeywebactioneventsnotificationschedulegroup `json:"blockingScheduleGroup,omitempty"`
		
		BlockingEmergencyScheduleGroup *Journeywebactioneventsnotificationemergencygroup `json:"blockingEmergencyScheduleGroup,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		ActionMap: o.ActionMap,
		
		ActionTarget: o.ActionTarget,
		
		BlockingReason: o.BlockingReason,
		
		BlockingActionMap: o.BlockingActionMap,
		
		BlockingAction: o.BlockingAction,
		
		BlockingFrequencyCapBehaviour: o.BlockingFrequencyCapBehaviour,
		
		BlockingPageUrlConditions: o.BlockingPageUrlConditions,
		
		BlockingScheduleGroup: o.BlockingScheduleGroup,
		
		BlockingEmergencyScheduleGroup: o.BlockingEmergencyScheduleGroup,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationblockedwebactionoffermessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationblockedwebactionoffermessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionMap, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ActionTarget, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if BlockingReason, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingReason"].(string); ok {
		o.BlockingReason = &BlockingReason
	}
    
	if BlockingActionMap, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingActionMap"].(map[string]interface{}); ok {
		BlockingActionMapString, _ := json.Marshal(BlockingActionMap)
		json.Unmarshal(BlockingActionMapString, &o.BlockingActionMap)
	}
	
	if BlockingAction, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingAction"].(map[string]interface{}); ok {
		BlockingActionString, _ := json.Marshal(BlockingAction)
		json.Unmarshal(BlockingActionString, &o.BlockingAction)
	}
	
	if BlockingFrequencyCapBehaviour, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingFrequencyCapBehaviour"].(string); ok {
		o.BlockingFrequencyCapBehaviour = &BlockingFrequencyCapBehaviour
	}
    
	if BlockingPageUrlConditions, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingPageUrlConditions"].([]interface{}); ok {
		BlockingPageUrlConditionsString, _ := json.Marshal(BlockingPageUrlConditions)
		json.Unmarshal(BlockingPageUrlConditionsString, &o.BlockingPageUrlConditions)
	}
	
	if BlockingScheduleGroup, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingScheduleGroup"].(map[string]interface{}); ok {
		BlockingScheduleGroupString, _ := json.Marshal(BlockingScheduleGroup)
		json.Unmarshal(BlockingScheduleGroupString, &o.BlockingScheduleGroup)
	}
	
	if BlockingEmergencyScheduleGroup, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingEmergencyScheduleGroup"].(map[string]interface{}); ok {
		BlockingEmergencyScheduleGroupString, _ := json.Marshal(BlockingEmergencyScheduleGroup)
		json.Unmarshal(BlockingEmergencyScheduleGroupString, &o.BlockingEmergencyScheduleGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
