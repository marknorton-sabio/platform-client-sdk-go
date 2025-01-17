package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatableimportjob - State information for an import job of rows to a datatable
type Datatableimportjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner - The PureCloud user who started the import job
	Owner *Addressableentityref `json:"owner,omitempty"`


	// Status - The status of the import job
	Status *string `json:"status,omitempty"`


	// DateCreated - The timestamp of when the import began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateCompleted - The timestamp of when the import stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// UploadURI - The URL of the location at which the caller can upload the file to be imported
	UploadURI *string `json:"uploadURI,omitempty"`


	// ImportMode - The indication of whether the processing should remove rows that don't appear in the import file
	ImportMode *string `json:"importMode,omitempty"`


	// ErrorInformation - Any error information, or null of the processing is not in an error state
	ErrorInformation *Errorbody `json:"errorInformation,omitempty"`


	// CountRecordsUpdated - The current count of the number of records processed
	CountRecordsUpdated *int `json:"countRecordsUpdated,omitempty"`


	// CountRecordsDeleted - The current count of the number of records deleted
	CountRecordsDeleted *int `json:"countRecordsDeleted,omitempty"`


	// CountRecordsFailed - The current count of the number of records that failed to import
	CountRecordsFailed *int `json:"countRecordsFailed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Datatableimportjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatableimportjob
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Owner *Addressableentityref `json:"owner,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		UploadURI *string `json:"uploadURI,omitempty"`
		
		ImportMode *string `json:"importMode,omitempty"`
		
		ErrorInformation *Errorbody `json:"errorInformation,omitempty"`
		
		CountRecordsUpdated *int `json:"countRecordsUpdated,omitempty"`
		
		CountRecordsDeleted *int `json:"countRecordsDeleted,omitempty"`
		
		CountRecordsFailed *int `json:"countRecordsFailed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Owner: o.Owner,
		
		Status: o.Status,
		
		DateCreated: DateCreated,
		
		DateCompleted: DateCompleted,
		
		UploadURI: o.UploadURI,
		
		ImportMode: o.ImportMode,
		
		ErrorInformation: o.ErrorInformation,
		
		CountRecordsUpdated: o.CountRecordsUpdated,
		
		CountRecordsDeleted: o.CountRecordsDeleted,
		
		CountRecordsFailed: o.CountRecordsFailed,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Datatableimportjob) UnmarshalJSON(b []byte) error {
	var DatatableimportjobMap map[string]interface{}
	err := json.Unmarshal(b, &DatatableimportjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DatatableimportjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DatatableimportjobMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Owner, ok := DatatableimportjobMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if Status, ok := DatatableimportjobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateCreatedString, ok := DatatableimportjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateCompletedString, ok := DatatableimportjobMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if UploadURI, ok := DatatableimportjobMap["uploadURI"].(string); ok {
		o.UploadURI = &UploadURI
	}
    
	if ImportMode, ok := DatatableimportjobMap["importMode"].(string); ok {
		o.ImportMode = &ImportMode
	}
    
	if ErrorInformation, ok := DatatableimportjobMap["errorInformation"].(map[string]interface{}); ok {
		ErrorInformationString, _ := json.Marshal(ErrorInformation)
		json.Unmarshal(ErrorInformationString, &o.ErrorInformation)
	}
	
	if CountRecordsUpdated, ok := DatatableimportjobMap["countRecordsUpdated"].(float64); ok {
		CountRecordsUpdatedInt := int(CountRecordsUpdated)
		o.CountRecordsUpdated = &CountRecordsUpdatedInt
	}
	
	if CountRecordsDeleted, ok := DatatableimportjobMap["countRecordsDeleted"].(float64); ok {
		CountRecordsDeletedInt := int(CountRecordsDeleted)
		o.CountRecordsDeleted = &CountRecordsDeletedInt
	}
	
	if CountRecordsFailed, ok := DatatableimportjobMap["countRecordsFailed"].(float64); ok {
		CountRecordsFailedInt := int(CountRecordsFailed)
		o.CountRecordsFailed = &CountRecordsFailedInt
	}
	
	if SelfUri, ok := DatatableimportjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Datatableimportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
