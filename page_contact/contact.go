package page_contact

import (
	. "Web/main_definitions"
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

type ContactWebPage struct {
	*PageData
	Formatting      string
	HomePage        string
	Message         string
	DefaultMessage  string
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

// ------------------------------------------- Public ------------------------------------------- //

// Initializes page
func (p *ContactWebPage) Init(baseData PageData) WebPageInterface {
	p.PageData = NewWebPage(baseData, "contact", "contact/", p.Handler)
	p.DefaultMessage = "Character Limit: 1000"
	p.Message = p.DefaultMessage
	p.Input = UserInput{}
	p.Formatting = p.UrlStaticFolder + "formatting.css"
	return p
}

// Expose common data fields
func (p *ContactWebPage) Data() *PageData {
	return p.PageData
}

// Implements page's behavior
func (p *ContactWebPage) Handler(w http.ResponseWriter, r *http.Request) {

	if p.HomePage == "" {
		p.HomePage = (*p.PageDict)["index"].Data().UrlExtension
	}

	// Get captcha data
	p.CaptchaLocation = (*p.PageDict)["captcha"].Data().UrlExtension

	captchaValue := GetData((*p.PageDict)["captcha"], "CaptchaCode", StringTypeArray)
	captchaText := captchaValue.(string)

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

	// Get user input data from html form
	input := UserInput{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		Captcha: r.FormValue("captcha"),
	}

	if len(input.Email) == 0 && len(input.Subject) == 0 && len(input.Message) == 0 && len(input.Captcha) == 0 {
		return
	} else if ok, msg := isValidInput(input, captchaText); !ok {
		p.Message = msg
		p.Input = input
	} else {
		p.Success = true
		generateAndSendEmail(input)
	}

	t.Execute(w, *p)

	if p.Success == true {
		// Reset data to default values
		p.Message = p.DefaultMessage
		p.Success = false
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
