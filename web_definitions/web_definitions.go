package web_definitions

import (
	// "github.com/gorilla/mux"
	// "github.com/gorilla/websocket"
	// "fmt"
	// "log"
	"net/http"
	// "template"
	"errors"
)

type WebPage struct {
	Name         string
	UrlExtension string
	// RootFolder   string
	SelfFolder string
	HtmlFile   string
	PageDict   *map[string]*WebPage
	Data       interface{}
	Handler    interface{}
}

func NewWebPage(pageName, urlExtension string, pageDict *map[string]*WebPage, handleFunc interface{}) *WebPage {
	p := WebPage{}
	p.Name = pageName
	p.UrlExtension = urlExtension
	p.SelfFolder = "template_" + p.Name + "\\"
	p.HtmlFile = p.SelfFolder + p.Name + ".html"
	p.PageDict = pageDict

	// Ensure that handleFunc is one of the two legal types
	_, okHandler := handleFunc.(http.Handler)
	_, okHandlerFunc := handleFunc.(func(http.ResponseWriter, *http.Request))
	if !(okHandler || okHandlerFunc) {
		err := errors.New("Invalid type for handleFunc")
		panic(err)
	}
	p.Handler = handleFunc

	return &p
}

func GetAllPages(pages map[string]func(*map[string]*WebPage) *WebPage) *map[string]*WebPage {

	pageDict := map[string]*WebPage{}
	for name, initFunction := range pages {
		pageDict[name] = initFunction(&pageDict)
	}
	return &pageDict
}

func GetAllPageUrls(pages *map[string]*WebPage) map[string]string {

	pageUrls := map[string]string{}
	for _, webPage := range *pages {
		pageUrls[webPage.Name] = webPage.UrlExtension
	}
	return pageUrls
}
