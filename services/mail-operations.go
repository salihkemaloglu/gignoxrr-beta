package helper 

import (
	"bytes"
	"html/template"
	"fmt"
	gomail "gopkg.in/gomail.v2"
)

func  SendUserRegisterConfirmationMail(userMail string,mailType string) string {
	
	mailTypePath:="app-root/mail-templates/"+mailType

	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return fmt.Sprintf("Mail template parse error: %v",err.Error())
	}

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, "test"); err != nil {
		return fmt.Sprintf("Mail template execute byte error: %v",err.Error())
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userMail)
	mail.SetHeader("Subject", "Confirm your account on Gignox")
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)

	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return fmt.Sprintf("Mail send  error: %v",err.Error())
	}
	return "ok"
}

