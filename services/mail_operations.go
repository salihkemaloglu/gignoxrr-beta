package helper 

import (
	"time"
	"bytes"
	"fmt"
	"html/template"
	gomail "gopkg.in/gomail.v2"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repository"
)
type UserRegisterData struct {
    WelcomeToGignax string
    ThanksForSigup string
    MustEnterVerrificationCode string
	VerificationCode string
	YourVerificationCode string
	OneUseCode string
}
func  SendUserRegisterConfirmationMail(userMail string,lang string,verificationCode string) string {
	
	mailTypePath:="app_root/mail_templates/user_register_confirmation.html"

	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return fmt.Sprintf("Mail template parse error: %v",err.Error())
	}
	wd := UserRegisterData{
        WelcomeToGignax: Translate(lang,"Welcome_To_Gignox"),
        ThanksForSigup: Translate(lang,"Thanks_For_Sigup"),
        MustEnterVerrificationCode: Translate(lang,"Must_Enter_Verrification_Code"),
        YourVerificationCode: Translate(lang,"Your_Verification_Code"),
        OneUseCode: Translate(lang,"One_Use_Code"),
        VerificationCode: verificationCode,
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return fmt.Sprintf("Mail template execute byte error: %v",err.Error())
	}
	t := time.Now().UTC()
	userTemporaryInformation := db.UserTemporaryInformation {
		Email: userMail,
		RegisterVerificationCode:verificationCode,
		ForgotPasswordVerificationCode:"",
		RegisterVerificationCodeCreateDate: t.Format("2006-01-02 15:04:05"),
		ForgotPasswordVerificationCodeCreateDate: t.Format("2006-01-02 15:04:05"),
		IsCodeUsed: false,
    	IsCodeExpired: false, 
	}
	var userOp repo.UserTemporaryInformationRepository=userTemporaryInformation
	if dbResp := userOp.Insert(); dbResp != nil {
		return fmt.Sprintf(Translate(lang,"User_Temporary_Indormation_Insert_Error")+" :%v",dbResp.Error())
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userMail)
	mail.SetHeader("Subject", Translate(lang,"User_Register_Mail_Subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return fmt.Sprintf("Mail send  error: %v",err.Error())
	}
	return "ok"
}

type UserForgotPasswordData struct {
    VerificationCodeTitle string
    ReceivedPasswordChangeRequest string
    ViaEmailAddress string
	DontShareVerificationCode string
	VerificationCode string
	YourVerificationCode string 
	OneUseCode string
	Email string
}

func  SendUserForgotPasswordVerificationMail(userMail string,lang string,verificationCode string) string {
	
	mailTypePath:="app_root/mail_templates/user_forgot_password.html"

	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return fmt.Sprintf("Mail template parse error: %v",err.Error())
	}
	wd := UserForgotPasswordData{
        VerificationCodeTitle: Translate(lang,"Verification_Code_Title"),
        ReceivedPasswordChangeRequest: Translate(lang,"Received_Password_Change_Request"),
        ViaEmailAddress: Translate(lang,"Via_Email_Address"),
		DontShareVerificationCode: Translate(lang,"Dont_Share_Verification_Code"),
		YourVerificationCode: Translate(lang,"Your_Verification_Code"),
        OneUseCode: Translate(lang,"One_Use_Code"),
		VerificationCode: verificationCode,
		Email:userMail,
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return fmt.Sprintf("Mail template execute byte error: %v",err.Error())
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userMail)
	mail.SetHeader("Subject", Translate(lang,"User_Forgot_Mail_Subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return fmt.Sprintf("Mail send  error: %v",err.Error())
	}
	return "ok"
}
