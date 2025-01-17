package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historylisting
type Historylisting struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Client
	Client *Domainentityref `json:"client,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Detail `json:"errorDetails,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`


	// ActionName - Action name
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus - Action status
	ActionStatus *string `json:"actionStatus,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// System
	System *bool `json:"system,omitempty"`


	// Started - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Started *time.Time `json:"started,omitempty"`


	// Completed - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Completed *time.Time `json:"completed,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Historyentry `json:"entities,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Historylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historylisting
	
	Started := new(string)
	if o.Started != nil {
		
		*Started = timeutil.Strftime(o.Started, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Started = nil
	}
	
	Completed := new(string)
	if o.Completed != nil {
		
		*Completed = timeutil.Strftime(o.Completed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Completed = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorDetails *[]Detail `json:"errorDetails,omitempty"`
		
		ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		System *bool `json:"system,omitempty"`
		
		Started *string `json:"started,omitempty"`
		
		Completed *string `json:"completed,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		Entities *[]Historyentry `json:"entities,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Complete: o.Complete,
		
		User: o.User,
		
		Client: o.Client,
		
		ErrorMessage: o.ErrorMessage,
		
		ErrorCode: o.ErrorCode,
		
		ErrorDetails: o.ErrorDetails,
		
		ErrorMessageParams: o.ErrorMessageParams,
		
		ActionName: o.ActionName,
		
		ActionStatus: o.ActionStatus,
		
		Name: o.Name,
		
		Description: o.Description,
		
		System: o.System,
		
		Started: Started,
		
		Completed: Completed,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		Entities: o.Entities,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Historylisting) UnmarshalJSON(b []byte) error {
	var HistorylistingMap map[string]interface{}
	err := json.Unmarshal(b, &HistorylistingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := HistorylistingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Complete, ok := HistorylistingMap["complete"].(bool); ok {
		o.Complete = &Complete
	}
    
	if User, ok := HistorylistingMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := HistorylistingMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if ErrorMessage, ok := HistorylistingMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if ErrorCode, ok := HistorylistingMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorDetails, ok := HistorylistingMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}
	
	if ErrorMessageParams, ok := HistorylistingMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}
	
	if ActionName, ok := HistorylistingMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
    
	if ActionStatus, ok := HistorylistingMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}
    
	if Name, ok := HistorylistingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := HistorylistingMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if System, ok := HistorylistingMap["system"].(bool); ok {
		o.System = &System
	}
    
	if startedString, ok := HistorylistingMap["started"].(string); ok {
		Started, _ := time.Parse("2006-01-02T15:04:05.999999Z", startedString)
		o.Started = &Started
	}
	
	if completedString, ok := HistorylistingMap["completed"].(string); ok {
		Completed, _ := time.Parse("2006-01-02T15:04:05.999999Z", completedString)
		o.Completed = &Completed
	}
	
	if PageSize, ok := HistorylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := HistorylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := HistorylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if Entities, ok := HistorylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageCount, ok := HistorylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
