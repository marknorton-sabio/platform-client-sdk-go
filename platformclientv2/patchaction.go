package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchaction
type Patchaction struct { 
	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`


	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`


	// ActionTargetId - Action target ID.
	ActionTargetId *string `json:"actionTargetId,omitempty"`


	// IsPacingEnabled - Whether this action should be throttled.
	IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`


	// Props - Additional properties.
	Props *Patchactionproperties `json:"props,omitempty"`


	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`


	// WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
	WebMessagingOfferFields *Patchwebmessagingofferfields `json:"webMessagingOfferFields,omitempty"`


	// OpenActionFields - Admin-configurable fields of an open action.
	OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`

}

func (o *Patchaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchaction
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		ActionTargetId *string `json:"actionTargetId,omitempty"`
		
		IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`
		
		Props *Patchactionproperties `json:"props,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Patchwebmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		
		OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
		*Alias
	}{ 
		MediaType: o.MediaType,
		
		ActionTemplate: o.ActionTemplate,
		
		ActionTargetId: o.ActionTargetId,
		
		IsPacingEnabled: o.IsPacingEnabled,
		
		Props: o.Props,
		
		ArchitectFlowFields: o.ArchitectFlowFields,
		
		WebMessagingOfferFields: o.WebMessagingOfferFields,
		
		OpenActionFields: o.OpenActionFields,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchaction) UnmarshalJSON(b []byte) error {
	var PatchactionMap map[string]interface{}
	err := json.Unmarshal(b, &PatchactionMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := PatchactionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ActionTemplate, ok := PatchactionMap["actionTemplate"].(map[string]interface{}); ok {
		ActionTemplateString, _ := json.Marshal(ActionTemplate)
		json.Unmarshal(ActionTemplateString, &o.ActionTemplate)
	}
	
	if ActionTargetId, ok := PatchactionMap["actionTargetId"].(string); ok {
		o.ActionTargetId = &ActionTargetId
	}
    
	if IsPacingEnabled, ok := PatchactionMap["isPacingEnabled"].(bool); ok {
		o.IsPacingEnabled = &IsPacingEnabled
	}
    
	if Props, ok := PatchactionMap["props"].(map[string]interface{}); ok {
		PropsString, _ := json.Marshal(Props)
		json.Unmarshal(PropsString, &o.Props)
	}
	
	if ArchitectFlowFields, ok := PatchactionMap["architectFlowFields"].(map[string]interface{}); ok {
		ArchitectFlowFieldsString, _ := json.Marshal(ArchitectFlowFields)
		json.Unmarshal(ArchitectFlowFieldsString, &o.ArchitectFlowFields)
	}
	
	if WebMessagingOfferFields, ok := PatchactionMap["webMessagingOfferFields"].(map[string]interface{}); ok {
		WebMessagingOfferFieldsString, _ := json.Marshal(WebMessagingOfferFields)
		json.Unmarshal(WebMessagingOfferFieldsString, &o.WebMessagingOfferFields)
	}
	
	if OpenActionFields, ok := PatchactionMap["openActionFields"].(map[string]interface{}); ok {
		OpenActionFieldsString, _ := json.Marshal(OpenActionFields)
		json.Unmarshal(OpenActionFieldsString, &o.OpenActionFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
