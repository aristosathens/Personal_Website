package template_contact

import (
	"gopkg.in/gomail.v2"
	// "github.com/steambap/captcha"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	. "web_definitions"
)

// ------------------------------------------- Definitions ------------------------------------------- //

var page *WebPage
var pageData PageData
var defaultMessage = "Character Limit: 1000"

type PageData struct {
	Message  string
	Input    UserInput
	Success  bool
	PageUrls map[string]string
}

type UserInput struct {
	Email   string
	Subject string
	Message string
}

// var defaultInput = UserInput{
// 	Email:   "",
// 	Subject: "",
// 	Message: "",
// }

// ------------------------------------------- Public ------------------------------------------- //

func Init(pageDict *map[string]*WebPage) *WebPage {
	page = NewWebPage("contact", "contact", pageDict, Handler)
	pageData = PageData{defaultMessage, UserInput{}, false, nil}
	return page
}

func Handler(w http.ResponseWriter, r *http.Request) {

	// If this is the first time, initialize pageUrls. This can only occur after all WebPage structs have finished initializing
	if pageData.PageUrls == nil {
		pageData.PageUrls = GetAllPageUrls(page.PageDict)
	}
	pageData.Message = defaultMessage
	pageData.Input = UserInput{}

	// func recursiveHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(page.HtmlFile) //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, pageData) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {              // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	input := UserInput{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	if input.Email == "" && input.Subject == "" && input.Message == "" {
		return
	}

	fmt.Println(input)

	if ok, msg := isValidInput(input); !ok {
		pageData.Message = msg
		pageData.Input = input
		err = t.Execute(w, pageData) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil {              // if there is an error
			log.Print("template executing error: ", err) //log it
		}
		return
	}

	generateAndSendEmail(input)
	t.Execute(w, struct{ Success bool }{true})
	// pageData.Message = defaultMessage
}

// ------------------------------------------- Private ------------------------------------------- //

// Populates the fields of an email and sends it to me from a dummy account
func generateAndSendEmail(input UserInput) {
	m := gomail.NewMessage()
	m.SetHeader("From", "Aristos.Website@gmail.com")
	m.SetHeader("To", "aristos.a.athens@gmail.com")
	m.SetHeader("Subject", input.Subject)
	m.SetBody("text/html", input.Message)

	d := gomail.NewDialer("smtp.gmail.com", 587, "Aristos.Website", "VerySecurePassword")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func isValidInput(input UserInput) (bool, string) {

	if !isValidEmail(input.Email) {
		return false, "Invalid email."
	}
	if len(input.Message) > 1000 {
		return false, "Please limit messages to 1000 characters or less."
	}
	if len(input.Subject) > 75 {
		return false, "Please limit subject to 75 characters or less."
	}
	return true, ""
}

// Checks if email is at least well formed (if not a real email)
func isValidEmail(input string) bool {

	index := strings.Index(input, "@")
	if index < 1 {
		return false
	}

	index = strings.Index(input[index:], ".")
	if index < 1 {
		return false
	}

	return true
}
