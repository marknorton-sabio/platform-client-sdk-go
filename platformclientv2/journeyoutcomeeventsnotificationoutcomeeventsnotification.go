package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcomeeventsnotification
type Journeyoutcomeeventsnotificationoutcomeeventsnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// ExternalContact
	ExternalContact *Journeyoutcomeeventsnotificationexternalcontact `json:"externalContact,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CustomerId
	CustomerId *string `json:"customerId,omitempty"`


	// CustomerIdType
	CustomerIdType *string `json:"customerIdType,omitempty"`


	// Session
	Session *Journeyoutcomeeventsnotificationsession `json:"session,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`


	// OutcomeAchievedEvent
	OutcomeAchievedEvent *Journeyoutcomeeventsnotificationoutcomeachievedmessage `json:"outcomeAchievedEvent,omitempty"`


	// OutcomeAttributionEventMessage
	OutcomeAttributionEventMessage *Journeyoutcomeeventsnotificationoutcomeattributionmessage `json:"outcomeAttributionEventMessage,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationoutcomeeventsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationoutcomeeventsnotification
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		ExternalContact *Journeyoutcomeeventsnotificationexternalcontact `json:"externalContact,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		Session *Journeyoutcomeeventsnotificationsession `json:"session,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		OutcomeAchievedEvent *Journeyoutcomeeventsnotificationoutcomeachievedmessage `json:"outcomeAchievedEvent,omitempty"`
		
		OutcomeAttributionEventMessage *Journeyoutcomeeventsnotificationoutcomeattributionmessage `json:"outcomeAttributionEventMessage,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		CorrelationId: o.CorrelationId,
		
		ExternalContact: o.ExternalContact,
		
		CreatedDate: CreatedDate,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		Session: o.Session,
		
		EventType: o.EventType,
		
		OutcomeAchievedEvent: o.OutcomeAchievedEvent,
		
		OutcomeAttributionEventMessage: o.OutcomeAttributionEventMessage,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcomeeventsnotification) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CorrelationId, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if ExternalContact, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if createdDateString, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CustomerId, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if Session, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["session"].(map[string]interface{}); ok {
		SessionString, _ := json.Marshal(Session)
		json.Unmarshal(SessionString, &o.Session)
	}
	
	if EventType, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if OutcomeAchievedEvent, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["outcomeAchievedEvent"].(map[string]interface{}); ok {
		OutcomeAchievedEventString, _ := json.Marshal(OutcomeAchievedEvent)
		json.Unmarshal(OutcomeAchievedEventString, &o.OutcomeAchievedEvent)
	}
	
	if OutcomeAttributionEventMessage, ok := JourneyoutcomeeventsnotificationoutcomeeventsnotificationMap["outcomeAttributionEventMessage"].(map[string]interface{}); ok {
		OutcomeAttributionEventMessageString, _ := json.Marshal(OutcomeAttributionEventMessage)
		json.Unmarshal(OutcomeAttributionEventMessageString, &o.OutcomeAttributionEventMessage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcomeeventsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
