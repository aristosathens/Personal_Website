package template_index

import (
	. "Web/web_definitions"
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

var page *IndexWebPage

type IndexWebPage struct {
	*PageData
	Formatting     string
	AristosPicture string
	ResumePage     string
	ContactPage    string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initiates page
func (p *IndexWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("index", "home/", localRootFolder, pageDict, IndexWebPageHandler)
	p.AristosPicture = p.UrlStaticFolder + "Aristos_Headshot.jpg"
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	page = p
	return p
}

// Implements page's behavior
func IndexWebPageHandler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, get data from other pages
	if page.ResumePage == "" {
		page.ResumePage = GetData((*page.PageDict)["resume"], "UrlExtension", StringTypeArray).(string)
		page.ContactPage = GetData((*page.PageDict)["contact"], "UrlExtension", StringTypeArray).(string)
	}

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
