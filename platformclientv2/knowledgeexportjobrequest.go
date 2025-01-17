package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeexportjobrequest
type Knowledgeexportjobrequest struct { 
	// ExportFilter - What to export.
	ExportFilter *Knowledgeexportjobfilter `json:"exportFilter,omitempty"`


	// FileType - File type of the document
	FileType *string `json:"fileType,omitempty"`

}

func (o *Knowledgeexportjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeexportjobrequest
	
	return json.Marshal(&struct { 
		ExportFilter *Knowledgeexportjobfilter `json:"exportFilter,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		*Alias
	}{ 
		ExportFilter: o.ExportFilter,
		
		FileType: o.FileType,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeexportjobrequest) UnmarshalJSON(b []byte) error {
	var KnowledgeexportjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeexportjobrequestMap)
	if err != nil {
		return err
	}
	
	if ExportFilter, ok := KnowledgeexportjobrequestMap["exportFilter"].(map[string]interface{}); ok {
		ExportFilterString, _ := json.Marshal(ExportFilter)
		json.Unmarshal(ExportFilterString, &o.ExportFilter)
	}
	
	if FileType, ok := KnowledgeexportjobrequestMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
