package page_resume

import (
	. "Web/main_definitions"
	"html/template"
	"log"
	"net/http"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// ResumeWebPage embeds the *WebPage type
// ResumeWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type ResumeWebPage struct {
	*PageData
	Formatting string
	Resume     string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *ResumeWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "resume", "resume/", p.Handler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	p.Resume = "https://drive.google.com/file/d/1geuuG-Qh_QUTOzkkgZ_z7nsa_SVhwv5f/preview"
	return p
}

// Expose common data fields
func (p *ResumeWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior
func (p *ResumeWebPage) Handler(w http.ResponseWriter, r *http.Request) {

	// Create Golang http template from html file
	t, err := template.ParseFiles(p.LocalHtmlFile)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Pass in the page's data and execute the template
	err = t.Execute(w, *p)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
