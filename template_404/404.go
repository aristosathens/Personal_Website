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

var page *FourZeroFourWebPage

type FourZeroFourWebPage struct {
	*PageData
	Formatting string
	HomePage   string
	Image      string
	Message    string
}

// ------------------------------------------- Public ------------------------------------------- //

// Initiates page
func (p *FourZeroFourWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("404", "", localRootFolder, pageDict, FourZeroFourWebPageWebPageHandler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	p.Image = p.UrlStaticFolder + "Sad_Otter.jpg"
	p.Message = "This page does not exist!"
	page = p
	return p
}

// Implements page's behavior. Writes status header, message, and link back to home
func FourZeroFourWebPageWebPageHandler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, get data from home page
	if page.HomePage == "" {
		page.HomePage = GetData((*page.PageDict)["index"], "UrlExtension", StringTypeArray).(string)
	}

	w.WriteHeader(http.StatusNotFound)

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
