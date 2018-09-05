package page_index

import (
	. "Web/main_definitions"
	"html/template"
	"log"
	"net/http"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// IndexWebPage embeds the *WebPage type
// IndexWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type IndexWebPage struct {
	*PageData
	Formatting     string
	AristosPicture string
	ResumePage     string
	ContactPage    string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *IndexWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "index", "home/", p.Handler)
	p.AristosPicture = p.UrlStaticFolder + "Aristos_Headshot.jpg"
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	return p
}

// Expose common data fields
func (p *IndexWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior
func (p *IndexWebPage) Handler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, get data from other pages
	if p.ResumePage == "" {
		p.ResumePage = (*p.PageDict)["resume"].Data().UrlExtension
		p.ContactPage = (*p.PageDict)["contact"].Data().UrlExtension
	}

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
