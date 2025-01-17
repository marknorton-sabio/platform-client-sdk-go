package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Deletableuserreference
type Deletableuserreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Delete - If marked true, the user will be removed an associated entity
	Delete *bool `json:"delete,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Deletableuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Deletableuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Delete: o.Delete,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Deletableuserreference) UnmarshalJSON(b []byte) error {
	var DeletableuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DeletableuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DeletableuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Delete, ok := DeletableuserreferenceMap["delete"].(bool); ok {
		o.Delete = &Delete
	}
    
	if SelfUri, ok := DeletableuserreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Deletableuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
