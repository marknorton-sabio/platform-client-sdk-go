package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentattachment - Attachment object.
type Contentattachment struct { 
	// Id - Provider specific ID for attachment. For example, a LINE sticker ID.
	Id *string `json:"id,omitempty"`


	// MediaType - The type of attachment this instance represents.
	MediaType *string `json:"mediaType,omitempty"`


	// Url - URL of the attachment.
	Url *string `json:"url,omitempty"`


	// Mime - Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
	Mime *string `json:"mime,omitempty"`


	// Text - Text associated with attachment such as an image caption.
	Text *string `json:"text,omitempty"`


	// Sha256 - Secure hash of the attachment content.
	Sha256 *string `json:"sha256,omitempty"`


	// Filename - Suggested file name for attachment.
	Filename *string `json:"filename,omitempty"`

}

func (o *Contentattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentattachment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Mime *string `json:"mime,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Sha256 *string `json:"sha256,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		Url: o.Url,
		
		Mime: o.Mime,
		
		Text: o.Text,
		
		Sha256: o.Sha256,
		
		Filename: o.Filename,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentattachment) UnmarshalJSON(b []byte) error {
	var ContentattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &ContentattachmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MediaType, ok := ContentattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Url, ok := ContentattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Mime, ok := ContentattachmentMap["mime"].(string); ok {
		o.Mime = &Mime
	}
    
	if Text, ok := ContentattachmentMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Sha256, ok := ContentattachmentMap["sha256"].(string); ok {
		o.Sha256 = &Sha256
	}
    
	if Filename, ok := ContentattachmentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
