package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationresponse
type Evaluationresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Conversation
	Conversation *Conversationreference `json:"conversation,omitempty"`


	// EvaluationForm - Evaluation form used for evaluation.
	EvaluationForm *Evaluationformresponse `json:"evaluationForm,omitempty"`


	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Calibration
	Calibration **Calibration `json:"calibration,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Answers
	Answers *Evaluationscoringset `json:"answers,omitempty"`


	// AgentHasRead
	AgentHasRead *bool `json:"agentHasRead,omitempty"`


	// ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`


	// AssignedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssignedDate *time.Time `json:"assignedDate,omitempty"`


	// ChangedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ChangedDate *time.Time `json:"changedDate,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// MediaType - List of different communication types used in conversation.
	MediaType *[]string `json:"mediaType,omitempty"`


	// Rescore - Is only true when evaluation is re-scored.
	Rescore *bool `json:"rescore,omitempty"`


	// ConversationDate - Date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationDate *time.Time `json:"conversationDate,omitempty"`


	// ConversationEndDate - End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationEndDate *time.Time `json:"conversationEndDate,omitempty"`


	// NeverRelease - Signifies if the evaluation is never to be released. This cannot be set true if release date is also set.
	NeverRelease *bool `json:"neverRelease,omitempty"`


	// ResourceId - Only used for email evaluations. Will be null for all other evaluations.
	ResourceId *string `json:"resourceId,omitempty"`


	// ResourceType - The type of resource. Only used for email evaluations. Will be null for evaluations on all other resources.
	ResourceType *string `json:"resourceType,omitempty"`


	// Redacted - Is only true when the user making the request does not have sufficient permissions to see evaluation
	Redacted *bool `json:"redacted,omitempty"`


	// IsScoringIndex
	IsScoringIndex *bool `json:"isScoringIndex,omitempty"`


	// AuthorizedActions - List of user authorized actions on evaluation. Possible values: edit, editScore, editAgentSignoff, delete, viewAudit
	AuthorizedActions *[]string `json:"authorizedActions,omitempty"`


	// HasAssistanceFailed - Is true when evaluation assistance didn't execute successfully
	HasAssistanceFailed *bool `json:"hasAssistanceFailed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Evaluationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationresponse
	
	ReleaseDate := new(string)
	if o.ReleaseDate != nil {
		
		*ReleaseDate = timeutil.Strftime(o.ReleaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReleaseDate = nil
	}
	
	AssignedDate := new(string)
	if o.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(o.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	ChangedDate := new(string)
	if o.ChangedDate != nil {
		
		*ChangedDate = timeutil.Strftime(o.ChangedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ChangedDate = nil
	}
	
	ConversationDate := new(string)
	if o.ConversationDate != nil {
		
		*ConversationDate = timeutil.Strftime(o.ConversationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationDate = nil
	}
	
	ConversationEndDate := new(string)
	if o.ConversationEndDate != nil {
		
		*ConversationEndDate = timeutil.Strftime(o.ConversationEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationformresponse `json:"evaluationForm,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Calibration **Calibration `json:"calibration,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		
		AgentHasRead *bool `json:"agentHasRead,omitempty"`
		
		ReleaseDate *string `json:"releaseDate,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		
		ChangedDate *string `json:"changedDate,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		MediaType *[]string `json:"mediaType,omitempty"`
		
		Rescore *bool `json:"rescore,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		ConversationEndDate *string `json:"conversationEndDate,omitempty"`
		
		NeverRelease *bool `json:"neverRelease,omitempty"`
		
		ResourceId *string `json:"resourceId,omitempty"`
		
		ResourceType *string `json:"resourceType,omitempty"`
		
		Redacted *bool `json:"redacted,omitempty"`
		
		IsScoringIndex *bool `json:"isScoringIndex,omitempty"`
		
		AuthorizedActions *[]string `json:"authorizedActions,omitempty"`
		
		HasAssistanceFailed *bool `json:"hasAssistanceFailed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		EvaluationForm: o.EvaluationForm,
		
		Evaluator: o.Evaluator,
		
		Agent: o.Agent,
		
		Calibration: o.Calibration,
		
		Status: o.Status,
		
		Answers: o.Answers,
		
		AgentHasRead: o.AgentHasRead,
		
		ReleaseDate: ReleaseDate,
		
		AssignedDate: AssignedDate,
		
		ChangedDate: ChangedDate,
		
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		Rescore: o.Rescore,
		
		ConversationDate: ConversationDate,
		
		ConversationEndDate: ConversationEndDate,
		
		NeverRelease: o.NeverRelease,
		
		ResourceId: o.ResourceId,
		
		ResourceType: o.ResourceType,
		
		Redacted: o.Redacted,
		
		IsScoringIndex: o.IsScoringIndex,
		
		AuthorizedActions: o.AuthorizedActions,
		
		HasAssistanceFailed: o.HasAssistanceFailed,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationresponse) UnmarshalJSON(b []byte) error {
	var EvaluationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluationresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := EvaluationresponseMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if EvaluationForm, ok := EvaluationresponseMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if Evaluator, ok := EvaluationresponseMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if Agent, ok := EvaluationresponseMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Calibration, ok := EvaluationresponseMap["calibration"].(map[string]interface{}); ok {
		CalibrationString, _ := json.Marshal(Calibration)
		json.Unmarshal(CalibrationString, &o.Calibration)
	}
	
	if Status, ok := EvaluationresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Answers, ok := EvaluationresponseMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if AgentHasRead, ok := EvaluationresponseMap["agentHasRead"].(bool); ok {
		o.AgentHasRead = &AgentHasRead
	}
    
	if releaseDateString, ok := EvaluationresponseMap["releaseDate"].(string); ok {
		ReleaseDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", releaseDateString)
		o.ReleaseDate = &ReleaseDate
	}
	
	if assignedDateString, ok := EvaluationresponseMap["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	
	if changedDateString, ok := EvaluationresponseMap["changedDate"].(string); ok {
		ChangedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", changedDateString)
		o.ChangedDate = &ChangedDate
	}
	
	if Queue, ok := EvaluationresponseMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := EvaluationresponseMap["mediaType"].([]interface{}); ok {
		MediaTypeString, _ := json.Marshal(MediaType)
		json.Unmarshal(MediaTypeString, &o.MediaType)
	}
	
	if Rescore, ok := EvaluationresponseMap["rescore"].(bool); ok {
		o.Rescore = &Rescore
	}
    
	if conversationDateString, ok := EvaluationresponseMap["conversationDate"].(string); ok {
		ConversationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationDateString)
		o.ConversationDate = &ConversationDate
	}
	
	if conversationEndDateString, ok := EvaluationresponseMap["conversationEndDate"].(string); ok {
		ConversationEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationEndDateString)
		o.ConversationEndDate = &ConversationEndDate
	}
	
	if NeverRelease, ok := EvaluationresponseMap["neverRelease"].(bool); ok {
		o.NeverRelease = &NeverRelease
	}
    
	if ResourceId, ok := EvaluationresponseMap["resourceId"].(string); ok {
		o.ResourceId = &ResourceId
	}
    
	if ResourceType, ok := EvaluationresponseMap["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
    
	if Redacted, ok := EvaluationresponseMap["redacted"].(bool); ok {
		o.Redacted = &Redacted
	}
    
	if IsScoringIndex, ok := EvaluationresponseMap["isScoringIndex"].(bool); ok {
		o.IsScoringIndex = &IsScoringIndex
	}
    
	if AuthorizedActions, ok := EvaluationresponseMap["authorizedActions"].([]interface{}); ok {
		AuthorizedActionsString, _ := json.Marshal(AuthorizedActions)
		json.Unmarshal(AuthorizedActionsString, &o.AuthorizedActions)
	}
	
	if HasAssistanceFailed, ok := EvaluationresponseMap["hasAssistanceFailed"].(bool); ok {
		o.HasAssistanceFailed = &HasAssistanceFailed
	}
    
	if SelfUri, ok := EvaluationresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
