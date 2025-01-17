package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventlog
type Eventlog struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorEntity
	ErrorEntity *Domainentityref `json:"errorEntity,omitempty"`


	// RelatedEntity
	RelatedEntity *Domainentityref `json:"relatedEntity,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Category
	Category *string `json:"category,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// EventMessage
	EventMessage *Eventmessage `json:"eventMessage,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Eventlog) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventlog
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ErrorEntity *Domainentityref `json:"errorEntity,omitempty"`
		
		RelatedEntity *Domainentityref `json:"relatedEntity,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		EventMessage *Eventmessage `json:"eventMessage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ErrorEntity: o.ErrorEntity,
		
		RelatedEntity: o.RelatedEntity,
		
		Timestamp: Timestamp,
		
		Level: o.Level,
		
		Category: o.Category,
		
		CorrelationId: o.CorrelationId,
		
		EventMessage: o.EventMessage,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventlog) UnmarshalJSON(b []byte) error {
	var EventlogMap map[string]interface{}
	err := json.Unmarshal(b, &EventlogMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EventlogMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EventlogMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ErrorEntity, ok := EventlogMap["errorEntity"].(map[string]interface{}); ok {
		ErrorEntityString, _ := json.Marshal(ErrorEntity)
		json.Unmarshal(ErrorEntityString, &o.ErrorEntity)
	}
	
	if RelatedEntity, ok := EventlogMap["relatedEntity"].(map[string]interface{}); ok {
		RelatedEntityString, _ := json.Marshal(RelatedEntity)
		json.Unmarshal(RelatedEntityString, &o.RelatedEntity)
	}
	
	if timestampString, ok := EventlogMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if Level, ok := EventlogMap["level"].(string); ok {
		o.Level = &Level
	}
    
	if Category, ok := EventlogMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if CorrelationId, ok := EventlogMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if EventMessage, ok := EventlogMap["eventMessage"].(map[string]interface{}); ok {
		EventMessageString, _ := json.Marshal(EventMessage)
		json.Unmarshal(EventMessageString, &o.EventMessage)
	}
	
	if SelfUri, ok := EventlogMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Eventlog) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
