package template_contact

import (
	. "Web/web_definitions"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// ContactWebPage embeds the *WebPage type
// ContactWebPage implements the WebPageInterface via its Init() function
// More details in web_definitions.go
//

var page *ContactWebPage

type ContactWebPage struct {
	*PageData
	Formatting      string
	HomePage        string
	Message         string
	Input           UserInput
	Success         bool
	CaptchaLocation string
}

type UserInput struct {
	Email   string
	Subject string
	Message string
	Captcha string
}

const defaultMessage = "Character Limit: 1000"

// ------------------------------------------- Public ------------------------------------------- //

func (p *ContactWebPage) Init(localRootFolder string, pageDict *map[string]WebPageInterface) WebPageInterface {
	p.PageData = NewWebPage("contact", "contact/", localRootFolder, pageDict, ContactWebPageHandler)
	p.Message = defaultMessage
	p.Input = UserInput{}
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	page = p
	return p
}

func ContactWebPageHandler(w http.ResponseWriter, r *http.Request) {

	if page.HomePage == "" {
		page.HomePage = GetData((*page.PageDict)["index"], "UrlExtension", StringTypeArray).(string)
	}

	// Get captcha data
	data := GetData((*page.PageDict)["captcha"], "UrlExtension", StringTypeArray)
	page.CaptchaLocation = data.(string)
	captchaValue := GetData((*page.PageDict)["captcha"], "CaptchaCode", StringTypeArray)
	captchaText := captchaValue.(string)

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

	// Get user input data from html form
	input := UserInput{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		Captcha: r.FormValue("captcha"),
	}

	if len(input.Email) == 0 && len(input.Subject) == 0 && len(input.Message) == 0 && len(input.Captcha) == 0 {
		// Default case. Do nothing
		return
	} else if ok, msg := isValidInput(input, captchaText); !ok {
		page.Message = msg
		page.Input = input
	} else {
		page.Success = true
	}

	// generateAndSendEmail(input)
	// page.Success = true
	t.Execute(w, *page)

	if page.Success == true {

		// Reset data to default values
		page.Message = defaultMessage
		page.Success = false
	}
}

// ------------------------------------------- Private ------------------------------------------- //

// Checks if user provided valid inputs for fields in html form
func isValidInput(input UserInput, captchaText string) (bool, string) {

	fmt.Println("Captcha text: " + captchaText)
	fmt.Println("User input: " + input.Captcha)

	if !isValidEmail(input.Email) {
		return false, "Invalid email."
	}
	if len(input.Subject) > 75 {
		return false, "Please limit subject to 75 characters or less."
	}
	if len(input.Subject) < 1 {
		return false, "Please include a subject."
	}
	if len(input.Message) > 1000 {
		return false, "Please limit messages to 1000 characters or less."
	}
	if len(input.Message) < 1 {
		return false, "Please include a message."
	}
	if input.Captcha != strings.TrimSpace(captchaText) {
		return false, "Incorrect captcha."
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
