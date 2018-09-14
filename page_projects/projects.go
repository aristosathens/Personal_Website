package page_projects

import (
	. "Web/main_definitions"

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// IndexWebPage embeds the *WebPage type
// IndexWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

type ProjectsWebPage struct {
	*PageData
	Formatting   string
	Projects     []Project
	TemplateHtml string
}

type Project struct {
	Name         string
	Formatting   string
	UrlExtension string
	TemplateHtml string
	Text         string
	Images       []string
	PDF          string
	Handler      func(http.ResponseWriter, *http.Request)
}

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *ProjectsWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "projects", "projects/", p.Handler)
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	p.TemplateHtml = p.LocalSelfFolder + "project_template.html"
	p.initProjectList()
	p.serveProjectPages()
	return p
}

// Expose common data fields
func (p *ProjectsWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior
func (p *ProjectsWebPage) Handler(w http.ResponseWriter, r *http.Request) {

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

// ------------------------------------------- Private ------------------------------------------- //

// Looks in static/projects/ for project folders, and populates the page's Projects field accordingly
func (p *ProjectsWebPage) initProjectList() {

	projectsLocal := p.LocalRootFolder + "static/projects/"
	projectsUrl := p.UrlRootFolder + "static/projects/"
	folders, err := ioutil.ReadDir(projectsLocal)
	if err != nil {
		return
	}

	// For each project subfolder, create new Project, and assign each file name to Project field
	p.Projects = []Project{}
	for _, folder := range folders {

		files, _ := ioutil.ReadDir(projectsLocal + folder.Name())
		if err != nil {
			continue
		}

		proj := Project{}
		proj.Name = folder.Name()
		proj.Formatting = p.Formatting
		proj.UrlExtension = p.UrlExtension + proj.Name + "/"
		proj.TemplateHtml = p.TemplateHtml
		proj.Handler = proj.projectPageHandler

		for _, file := range files {
			fName := file.Name()
			if endsWith(fName, []string{"txt"}) {
				proj.Text = projectsUrl + fName
			} else if endsWith(fName, []string{"pdf", "doc", "docx"}) {
				proj.PDF = projectsUrl + fName
			} else if endsWith(fName, []string{"png", "jpg", "jpeg", "gif", "tif"}) {
				proj.Images = append(proj.Images, projectsUrl+proj.Name+"/"+fName)
			}
		}
		p.Projects = append(p.Projects, proj)
	}
}

// Serve each project's page using the generic handler
func (p *ProjectsWebPage) serveProjectPages() {
	for _, project := range p.Projects {
		http.HandleFunc(project.UrlExtension, project.Handler)
	}
}

// Generic handler for displaying each project's page
func (proj *Project) projectPageHandler(w http.ResponseWriter, r *http.Request) {

	// Create Golang http template from html file
	t, err := template.ParseFiles(proj.TemplateHtml)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Pass in the page's data and execute the template
	err = t.Execute(w, *proj)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// ------------------------------------------- Utility ------------------------------------------- //

// Checks if string ends with any of the strings in endings
func endsWith(toCheck string, endings []string) bool {
	returnVal := false
	for _, ending := range endings {
		if strings.HasSuffix(toCheck, ending) {
			returnVal = true
		}
	}
	return returnVal
}