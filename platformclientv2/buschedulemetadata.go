package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulemetadata
type Buschedulemetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks spanned by this schedule
	WeekCount *int `json:"weekCount,omitempty"`


	// Description - The description of this schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether this schedule is published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast - The forecast used for this schedule, if applicable
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// GenerationResults - Generation result summary for this schedule, if applicable
	GenerationResults *Schedulegenerationresultsummary `json:"generationResults,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// ManagementUnits - High level per-management unit schedule metadata
	ManagementUnits *[]Bumanagementunitschedulesummary `json:"managementUnits,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Buschedulemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulemetadata
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		GenerationResults *Schedulegenerationresultsummary `json:"generationResults,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		ManagementUnits *[]Bumanagementunitschedulesummary `json:"managementUnits,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: o.WeekCount,
		
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		GenerationResults: o.GenerationResults,
		
		Metadata: o.Metadata,
		
		ManagementUnits: o.ManagementUnits,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulemetadata) UnmarshalJSON(b []byte) error {
	var BuschedulemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulemetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BuschedulemetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := BuschedulemetadataMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if WeekCount, ok := BuschedulemetadataMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Description, ok := BuschedulemetadataMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := BuschedulemetadataMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ShortTermForecast, ok := BuschedulemetadataMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if GenerationResults, ok := BuschedulemetadataMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if Metadata, ok := BuschedulemetadataMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if ManagementUnits, ok := BuschedulemetadataMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	
	if SelfUri, ok := BuschedulemetadataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
