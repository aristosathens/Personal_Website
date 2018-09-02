package template_resume

import (
	"html/template"
	"log"
	"net/http"
	. "web_definitions"
)

var page *WebPage
var pageUrls map[string]string

func Init(pageDict *map[string]*WebPage) *WebPage {
	page = NewWebPage("resume", "resume", pageDict, Handler)
	return page
}

func Handler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, initialize pageUrls. This can only occur after all WebPage structs have finished initializing
	if pageUrls == nil {
		pageUrls = GetAllPageUrls(page.PageDict)
	}
	pageUrls["resume"] = "static/Aristos_Resume.pdf"

	t, err := template.ParseFiles(page.HtmlFile) //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, pageUrls) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {              // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
