package page_about

import (
	. "Web/main_definitions"
	"html/template"
	"log"
	"net/http"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// AboutWebPage embeds the *WebPage type
// AboutWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type AboutWebPage struct {
	*PageData
	Formatting string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *AboutWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "about", "about/", p.Handler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	return p
}

// Expose common data fields
func (p *AboutWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior
func (p *AboutWebPage) Handler(w http.ResponseWriter, r *http.Request) {

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
