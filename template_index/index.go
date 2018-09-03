package template_index

import (
	. "Web/web_definitions"
	"html/template"
	"log"
	"net/http"
)

// ------------------------------------------- Definitions ------------------------------------------- //

//
// IndexWebPage embeds the *WebPage type
// IndexWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

var page *IndexWebPage

type IndexWebPage struct {
	*PageData
	Formatting     string
	PageUrls       map[string]string
	AristosPicture string
}

// ------------------------------------------- Public ------------------------------------------- //

func (p *IndexWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("index", "", localRootFolder, pageDict, IndexWebPageHandler)
	page = p
	return p
}

// func (p *IndexWebPage) GetPageData() *PageData {
// 	return p.PageData
// }

func IndexWebPageHandler(w http.ResponseWriter, r *http.Request) {

	// If this is the first call, initialize data. GetAllPageUrls() can only be called after all WebPage structs have finished initializing
	if page.PageUrls == nil {
		// page.PageUrls = GetAllPageUrls(page.PageDict)
		page.AristosPicture = page.UrlStaticFolder + "Aristos_Headshot.jpg"
		// page.Formatting = page.UrlStaticFolder + "formatting.css"
	}
	page.Formatting = page.UrlStaticFolder + "formatting.css"

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
