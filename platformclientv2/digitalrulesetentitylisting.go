package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitalrulesetentitylisting
type Digitalrulesetentitylisting struct { 
	// Entities
	Entities *[]Digitalruleset `json:"entities,omitempty"`


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

func (o *Digitalrulesetentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Digitalrulesetentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Digitalruleset `json:"entities,omitempty"`
		
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

func (o *Digitalrulesetentitylisting) UnmarshalJSON(b []byte) error {
	var DigitalrulesetentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &DigitalrulesetentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := DigitalrulesetentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := DigitalrulesetentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := DigitalrulesetentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := DigitalrulesetentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := DigitalrulesetentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if NextUri, ok := DigitalrulesetentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if LastUri, ok := DigitalrulesetentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if SelfUri, ok := DigitalrulesetentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PreviousUri, ok := DigitalrulesetentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if PageCount, ok := DigitalrulesetentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Digitalrulesetentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
