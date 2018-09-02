package template_formatting

import (
	"net/http"
	. "web_definitions"
)

var page *WebPage

func Init(pageDict *map[string]*WebPage) *WebPage {
	page = NewWebPage("formatting", "template_formatting/formatting.css", pageDict, Handler())
	return page
}

func Handler() http.Handler {
	cssFileServer := http.FileServer(http.Dir("template_formatting"))
	return http.StripPrefix("/template_formatting/", cssFileServer)
}
