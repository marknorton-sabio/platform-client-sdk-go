package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userrecording
type Userrecording struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userrecording
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ContentUri: o.ContentUri,
		
		Workspace: o.Workspace,
		
		CreatedBy: o.CreatedBy,
		
		Conversation: o.Conversation,
		
		ContentLength: o.ContentLength,
		
		DurationMilliseconds: o.DurationMilliseconds,
		
		Thumbnails: o.Thumbnails,
		
		Read: o.Read,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userrecording) UnmarshalJSON(b []byte) error {
	var UserrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &UserrecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserrecordingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserrecordingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := UserrecordingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := UserrecordingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ContentUri, ok := UserrecordingMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if Workspace, ok := UserrecordingMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := UserrecordingMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Conversation, ok := UserrecordingMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if ContentLength, ok := UserrecordingMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if DurationMilliseconds, ok := UserrecordingMap["durationMilliseconds"].(float64); ok {
		DurationMillisecondsInt := int(DurationMilliseconds)
		o.DurationMilliseconds = &DurationMillisecondsInt
	}
	
	if Thumbnails, ok := UserrecordingMap["thumbnails"].([]interface{}); ok {
		ThumbnailsString, _ := json.Marshal(Thumbnails)
		json.Unmarshal(ThumbnailsString, &o.Thumbnails)
	}
	
	if Read, ok := UserrecordingMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if SelfUri, ok := UserrecordingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
