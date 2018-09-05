package template_404

import (
	. "Web/web_definitions"
	"html/template"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	// "strings"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// FourZeroFourWebPage embeds the *WebPage type
// FourZeroFourWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type FourZeroFourWebPage struct {
	*PageData
	Formatting string
	HomePage   string
	Image      string
	Message    string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *FourZeroFourWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "404", "", p.Handler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	p.Image = p.UrlStaticFolder + "Sad_Otter.jpg"
	p.Message = "This page does not exist!"
	return p
}

// Expose common data fields
func (p *FourZeroFourWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior. Writes status header, message, and link back to home
func (p *FourZeroFourWebPage) Handler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, get data from home page
	if p.HomePage == "" {
		p.HomePage = (*p.PageDict)["index"].Data().UrlExtension
	}

	w.WriteHeader(http.StatusNotFound)

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
