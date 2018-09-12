package page_captcha

import (
	. "Web/main_definitions"
	"github.com/steambap/captcha"
	"log"
	"net/http"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// CaptchaWebPage embeds the *WebPage type
// CaptchaWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type CaptchaWebPage struct {
	*PageData
	CaptchaCode string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *CaptchaWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "captcha", "captcha/", p.Handler)
	return p
}

// Expose common data fields
func (p *CaptchaWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior. Generate new captcha and write png to ResponseWriter
func (p *CaptchaWebPage) Handler(w http.ResponseWriter, r *http.Request) {
	img, err := captcha.New(250, 250)
	if err != nil {
		log.Print("Captcha creation error: ", err)
		return
	}
	p.CaptchaCode = img.Text
	img.WriteImage(w)
}
