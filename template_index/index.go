package template_index

import (
	"html/template"
	"log"
	"net/http"
	. "web_definitions"
)

// ------------------------------------------- Definitions ------------------------------------------- //

var pageUrls map[string]string
var page *WebPage

func Init(pageDict *map[string]*WebPage) *WebPage {
	page = NewWebPage("index", "", pageDict, Handler)
	return page
}

// ------------------------------------------- Public ------------------------------------------- //

func Handler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, initialize pageUrls. This can only occur after all WebPage structs have finished initializing
	if pageUrls == nil {
		pageUrls = GetAllPageUrls(page.PageDict)
	}
	pageUrls["Aristos_Picture"] = "static/Aristos_Headshot.jpg"

	t, err := template.ParseFiles(page.HtmlFile) //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, pageUrls) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
