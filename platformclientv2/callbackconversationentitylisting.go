package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackconversationentitylisting
type Callbackconversationentitylisting struct { 
	// Entities
	Entities *[]Callbackconversation `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Callbackconversationentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callbackconversationentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Callbackconversation `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		NextUri: o.NextUri,
		
		LastUri: o.LastUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Callbackconversationentitylisting) UnmarshalJSON(b []byte) error {
	var CallbackconversationentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackconversationentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := CallbackconversationentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := CallbackconversationentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := CallbackconversationentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := CallbackconversationentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := CallbackconversationentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if NextUri, ok := CallbackconversationentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if LastUri, ok := CallbackconversationentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if SelfUri, ok := CallbackconversationentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PreviousUri, ok := CallbackconversationentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if PageCount, ok := CallbackconversationentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackconversationentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
