package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentquestiongroupscore
type Assessmentquestiongroupscore struct { 
	// QuestionGroupId - The ID of the question group
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore - The total score for the questions
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore - The maximum total score for the questions
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA - True if this question group is marked NA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// TotalCriticalScore - The total score for the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// MaxTotalCriticalScore - The maximum total score for the critical questions
	MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`


	// TotalNonCriticalScore - The total score for the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// MaxTotalNonCriticalScore - The maximum total score for the non-critical questions
	MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`


	// TotalScoreUnweighted - The unweighted total score for this question group
	TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`


	// MaxTotalScoreUnweighted - The maximum unweighted total score for this question group
	MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`


	// TotalCriticalScoreUnweighted - The unweighted total score for the critical questions
	TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`


	// MaxTotalCriticalScoreUnweighted - The maximum unweighted total score for the critical questions
	MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`


	// TotalNonCriticalScoreUnweighted - The total unweighted score for the non-critical questions
	TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`


	// MaxTotalNonCriticalScoreUnweighted - The maximum unweighted total score for the non-critical questions
	MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`


	// QuestionScores - The individual question scores
	QuestionScores *[]Assessmentquestionscore `json:"questionScores,omitempty"`

}

func (o *Assessmentquestiongroupscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentquestiongroupscore
	
	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`
		
		TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`
		
		MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`
		
		TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`
		
		MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`
		
		TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`
		
		MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`
		
		QuestionScores *[]Assessmentquestionscore `json:"questionScores,omitempty"`
		*Alias
	}{ 
		QuestionGroupId: o.QuestionGroupId,
		
		TotalScore: o.TotalScore,
		
		MaxTotalScore: o.MaxTotalScore,
		
		MarkedNA: o.MarkedNA,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		MaxTotalCriticalScore: o.MaxTotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		MaxTotalNonCriticalScore: o.MaxTotalNonCriticalScore,
		
		TotalScoreUnweighted: o.TotalScoreUnweighted,
		
		MaxTotalScoreUnweighted: o.MaxTotalScoreUnweighted,
		
		TotalCriticalScoreUnweighted: o.TotalCriticalScoreUnweighted,
		
		MaxTotalCriticalScoreUnweighted: o.MaxTotalCriticalScoreUnweighted,
		
		TotalNonCriticalScoreUnweighted: o.TotalNonCriticalScoreUnweighted,
		
		MaxTotalNonCriticalScoreUnweighted: o.MaxTotalNonCriticalScoreUnweighted,
		
		QuestionScores: o.QuestionScores,
		Alias:    (*Alias)(o),
	})
}

func (o *Assessmentquestiongroupscore) UnmarshalJSON(b []byte) error {
	var AssessmentquestiongroupscoreMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentquestiongroupscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionGroupId, ok := AssessmentquestiongroupscoreMap["questionGroupId"].(string); ok {
		o.QuestionGroupId = &QuestionGroupId
	}
    
	if TotalScore, ok := AssessmentquestiongroupscoreMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if MaxTotalScore, ok := AssessmentquestiongroupscoreMap["maxTotalScore"].(float64); ok {
		MaxTotalScoreFloat32 := float32(MaxTotalScore)
		o.MaxTotalScore = &MaxTotalScoreFloat32
	}
	
	if MarkedNA, ok := AssessmentquestiongroupscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if TotalCriticalScore, ok := AssessmentquestiongroupscoreMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if MaxTotalCriticalScore, ok := AssessmentquestiongroupscoreMap["maxTotalCriticalScore"].(float64); ok {
		MaxTotalCriticalScoreFloat32 := float32(MaxTotalCriticalScore)
		o.MaxTotalCriticalScore = &MaxTotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := AssessmentquestiongroupscoreMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if MaxTotalNonCriticalScore, ok := AssessmentquestiongroupscoreMap["maxTotalNonCriticalScore"].(float64); ok {
		MaxTotalNonCriticalScoreFloat32 := float32(MaxTotalNonCriticalScore)
		o.MaxTotalNonCriticalScore = &MaxTotalNonCriticalScoreFloat32
	}
	
	if TotalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["totalScoreUnweighted"].(float64); ok {
		TotalScoreUnweightedFloat32 := float32(TotalScoreUnweighted)
		o.TotalScoreUnweighted = &TotalScoreUnweightedFloat32
	}
	
	if MaxTotalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["maxTotalScoreUnweighted"].(float64); ok {
		MaxTotalScoreUnweightedFloat32 := float32(MaxTotalScoreUnweighted)
		o.MaxTotalScoreUnweighted = &MaxTotalScoreUnweightedFloat32
	}
	
	if TotalCriticalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["totalCriticalScoreUnweighted"].(float64); ok {
		TotalCriticalScoreUnweightedFloat32 := float32(TotalCriticalScoreUnweighted)
		o.TotalCriticalScoreUnweighted = &TotalCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalCriticalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["maxTotalCriticalScoreUnweighted"].(float64); ok {
		MaxTotalCriticalScoreUnweightedFloat32 := float32(MaxTotalCriticalScoreUnweighted)
		o.MaxTotalCriticalScoreUnweighted = &MaxTotalCriticalScoreUnweightedFloat32
	}
	
	if TotalNonCriticalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["totalNonCriticalScoreUnweighted"].(float64); ok {
		TotalNonCriticalScoreUnweightedFloat32 := float32(TotalNonCriticalScoreUnweighted)
		o.TotalNonCriticalScoreUnweighted = &TotalNonCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalNonCriticalScoreUnweighted, ok := AssessmentquestiongroupscoreMap["maxTotalNonCriticalScoreUnweighted"].(float64); ok {
		MaxTotalNonCriticalScoreUnweightedFloat32 := float32(MaxTotalNonCriticalScoreUnweighted)
		o.MaxTotalNonCriticalScoreUnweighted = &MaxTotalNonCriticalScoreUnweightedFloat32
	}
	
	if QuestionScores, ok := AssessmentquestiongroupscoreMap["questionScores"].([]interface{}); ok {
		QuestionScoresString, _ := json.Marshal(QuestionScores)
		json.Unmarshal(QuestionScoresString, &o.QuestionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
