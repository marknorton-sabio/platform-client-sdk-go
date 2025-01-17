package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Segment
type Segment struct { 
	// StartTime - The timestamp when this segment began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The timestamp when this segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// VarType - The activity taking place for the participant in the segment.
	VarType *string `json:"type,omitempty"`


	// HowEnded - A description of the event that ended the segment.
	HowEnded *string `json:"howEnded,omitempty"`


	// DisconnectType - A description of the event that disconnected the segment
	DisconnectType *string `json:"disconnectType,omitempty"`

}

func (o *Segment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Segment
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		HowEnded *string `json:"howEnded,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		*Alias
	}{ 
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		VarType: o.VarType,
		
		HowEnded: o.HowEnded,
		
		DisconnectType: o.DisconnectType,
		Alias:    (*Alias)(o),
	})
}

func (o *Segment) UnmarshalJSON(b []byte) error {
	var SegmentMap map[string]interface{}
	err := json.Unmarshal(b, &SegmentMap)
	if err != nil {
		return err
	}
	
	if startTimeString, ok := SegmentMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := SegmentMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if VarType, ok := SegmentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if HowEnded, ok := SegmentMap["howEnded"].(string); ok {
		o.HowEnded = &HowEnded
	}
    
	if DisconnectType, ok := SegmentMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Segment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
