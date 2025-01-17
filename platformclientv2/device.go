package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Device
type Device struct { 
	// Category - Device category.
	Category *string `json:"category,omitempty"`


	// VarType - Device type (e.g. iPad, iPhone, Other).
	VarType *string `json:"type,omitempty"`


	// IsMobile - Flag that is true for mobile devices.
	IsMobile *bool `json:"isMobile,omitempty"`


	// ScreenHeight - Device's screen height.
	ScreenHeight *int `json:"screenHeight,omitempty"`


	// ScreenWidth - Device's screen width.
	ScreenWidth *int `json:"screenWidth,omitempty"`


	// Fingerprint - Fingerprint generated by looking at the individual device features.
	Fingerprint *string `json:"fingerprint,omitempty"`


	// OsFamily - Operating system family.
	OsFamily *string `json:"osFamily,omitempty"`


	// OsVersion - Operating system version.
	OsVersion *string `json:"osVersion,omitempty"`

}

func (o *Device) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Device
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		IsMobile *bool `json:"isMobile,omitempty"`
		
		ScreenHeight *int `json:"screenHeight,omitempty"`
		
		ScreenWidth *int `json:"screenWidth,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		OsFamily *string `json:"osFamily,omitempty"`
		
		OsVersion *string `json:"osVersion,omitempty"`
		*Alias
	}{ 
		Category: o.Category,
		
		VarType: o.VarType,
		
		IsMobile: o.IsMobile,
		
		ScreenHeight: o.ScreenHeight,
		
		ScreenWidth: o.ScreenWidth,
		
		Fingerprint: o.Fingerprint,
		
		OsFamily: o.OsFamily,
		
		OsVersion: o.OsVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Device) UnmarshalJSON(b []byte) error {
	var DeviceMap map[string]interface{}
	err := json.Unmarshal(b, &DeviceMap)
	if err != nil {
		return err
	}
	
	if Category, ok := DeviceMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if VarType, ok := DeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsMobile, ok := DeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
    
	if ScreenHeight, ok := DeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := DeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if Fingerprint, ok := DeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if OsFamily, ok := DeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
    
	if OsVersion, ok := DeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Device) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
