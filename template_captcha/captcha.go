package template_captcha

import (
	. "Web/web_definitions"
	"github.com/steambap/captcha"
	"log"
	"net/http"
)

// ------------------------------------------- Definitions ------------------------------------------- //

//
// CaptchaWebPage embeds the *WebPage type
// CaptchaWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

var page *CaptchaWebPage

type CaptchaWebPage struct {
	*PageData
	CaptchaCode string
}

// ------------------------------------------- Public ------------------------------------------- //

func (p *CaptchaWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("captcha", "captcha/", localRootFolder, pageDict, CaptchaWebPageHandler)
	page = p
	return p
}

// Implements page's behavior
func CaptchaWebPageHandler(w http.ResponseWriter, r *http.Request) {
	img, err := captcha.New(250, 75)
	if err != nil {
		log.Print("Captcha creation error: ", err)
		return
	}
	page.CaptchaCode = img.Text
	img.WriteImage(w)
}
