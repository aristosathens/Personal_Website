package template_resume

import (
	. "Web/web_definitions"
	"html/template"
	"log"
	"net/http"
)

// ------------------------------------------- Definitions ------------------------------------------- //

//
// ResumeWebPage embeds the *WebPage type
// ResumeWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

var page *ResumeWebPage

type ResumeWebPage struct {
	*PageData
	Formatting string
	Resume     string
}

// ------------------------------------------- Public ------------------------------------------- //

func (p *ResumeWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("resume", "resume/", localRootFolder, pageDict, ResumeWebPageHandler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	p.Resume = "https://drive.google.com/file/d/1geuuG-Qh_QUTOzkkgZ_z7nsa_SVhwv5f/preview"
	page = p
	return p
}

// Implements page's behavior
func ResumeWebPageHandler(w http.ResponseWriter, r *http.Request) {

	// Create Golang http template from html file
	t, err := template.ParseFiles(page.LocalHtmlFile)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Pass in the page's data and execute the template
	err = t.Execute(w, *page)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
