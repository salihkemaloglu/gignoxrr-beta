package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	inter "github.com/salihkemaloglu/gignoxrr-beta-001/interfaces"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gomail "gopkg.in/gomail.v2"
)

//UserRegisterData ...
type UserRegisterData struct {
	WelcomeToGignax                string
	ThanksForSigup                 string
	ClickVerrificationLink         string
	VerificationLink               string
	YourVerificationLink           string
	OnceVerifiedAccount            string
	UserVerifiedAccountInformation string
	OneUseLink                     string
	Email                          string
}

// UserForgotPasswordData ...
type UserForgotPasswordData struct {
	VerificationTokenTitle        string
	ReceivedPasswordChangeRequest string
	ViaEmailAddress               string
	DontShareVerificationLink     string
	ResetPasswordLink             string
	YourVerificationLink          string
	OneUseToken                   string
	Email                         string
}

//SendUserRegisterConfirmationMailService ...
func SendUserRegisterConfirmationMailService(userEmail string, username string, emailType string, verificationToken string, lang string) (*gigxRR.SendEmailResponse, error) {

	t := time.Now().UTC()
	userTemporaryInformation := repo.UserTemporaryInformation{
		Email:                                     userEmail,
		ForgotPasswordVerificationToken:           "",
		RegisterVerificationTokenCreateDate:       t.Format("2006-01-02 15:04:05"),
		ForgotPasswordVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		EmailType:                                 emailType,
		IsTokenUsed:                               false,
		IsTokenExpired:                            false,
	}
	// if there is verification code not used
	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository = &userTemporaryInformation
	if dbResp, dbErr := userTemporaryInformationOp.CheckVerificationTokenResentEmail(); dbErr == nil {
		userTemporaryInformation.ID = dbResp.ID
		userTemporaryInformation.RegisterVerificationToken = dbResp.RegisterVerificationToken
		dbResp.IsTokenUsed = false
		dbResp.IsTokenExpired = true
		if dbErr := userTemporaryInformationOp.Update(); dbErr != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintf(Translate(lang, "resent_register_email_update_database_error")+":%v", dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.RegisterVerificationToken = verificationToken
	mailTypePath := "app_root/mail_templates/user_register_confirmation.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_template_parse_error")+":%v", err.Error()),
		)
	}
	var verificationLink = "https://gignox.com/" + username + "/" + verificationToken
	wd := UserRegisterData{
		WelcomeToGignax:                Translate(lang, "welcome_to_gignox"),
		ThanksForSigup:                 Translate(lang, "thanks_for_signup"),
		ClickVerrificationLink:         Translate(lang, "click_verification_link"),
		YourVerificationLink:           Translate(lang, "your_verification_link"),
		OnceVerifiedAccount:            Translate(lang, "once_verified_user_account"),
		UserVerifiedAccountInformation: Translate(lang, "verified_user_account_information"),
		VerificationLink:               verificationLink,
		Email:                          userEmail,
	}

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_template_execute_error")+":%v", err.Error()),
		)
	}

	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "user_temporary_information_insert_error")+":%v", err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail)
	mail.SetHeader("Subject", Translate(lang, "user_register_mail_subject")) //bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)

	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_send_error")+":%v", err.Error()),
		)
	}
	return &gigxRR.SendEmailResponse{
		GeneralResponse: &gigxRR.GeneralResponse{
			Message:            emailType,
			IsEmailSuccess:     true,
			IsOperationSuccess: true,
		},
	}, nil
}

//SendUserForgotPasswordVerificationMailService ...
func SendUserForgotPasswordVerificationMailService(userEmail string, emailType string, verificationToken string, lang string) (*gigxRR.SendEmailResponse, error) {

	t := time.Now().UTC()
	userTemporaryInformation := repo.UserTemporaryInformation{
		Email:                               userEmail,
		RegisterVerificationToken:           "",
		RegisterVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		ForgotPasswordVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		EmailType:      emailType,
		IsTokenUsed:    false,
		IsTokenExpired: false,
	}
	// if there is verification code not used
	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository = &userTemporaryInformation
	if dbResp, dbErr := userTemporaryInformationOp.CheckVerificationTokenResentEmail(); dbErr == nil {
		userTemporaryInformation.ID = dbResp.ID
		userTemporaryInformation.ForgotPasswordVerificationToken = dbResp.ForgotPasswordVerificationToken
		userTemporaryInformation.IsTokenUsed = false
		userTemporaryInformation.IsTokenExpired = true
		if dbErr := userTemporaryInformationOp.Update(); dbErr != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintf(Translate(lang, "resent_forgot_password_email_update_database_error")+":%v", dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.ForgotPasswordVerificationToken = verificationToken
	mailTypePath := "app_root/mail_templates/user_forgot_password.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_template_parse_error")+":%v", err.Error()),
		)
	}
	var resetPasswordLink = "https://gignox.com/password_reset/" + verificationToken
	wd := UserForgotPasswordData{
		VerificationTokenTitle:        Translate(lang, "password_reset_token_title"),
		ReceivedPasswordChangeRequest: Translate(lang, "received_password_change_request"),
		ViaEmailAddress:               Translate(lang, "via_email_address"),
		DontShareVerificationLink:     Translate(lang, "dont_share_verification_token"),
		YourVerificationLink:          Translate(lang, "your_password_reset_link"),
		OneUseToken:                   Translate(lang, "one_use_link"),
		ResetPasswordLink:             resetPasswordLink,
		Email:                         userEmail,
	}

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_template_execute_error")+":%v", err.Error()),
		)
	}
	userTemporaryInformation.IsTokenUsed = false
	userTemporaryInformation.IsTokenExpired = false
	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "user_temporary_information_insert_error")+":%v", err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail)
	mail.SetHeader("Subject", Translate(lang, "user_forgot_mail_subject")) //bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)

	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang, "email_send_error")+":%v", err.Error()),
		)
	}
	return &gigxRR.SendEmailResponse{
		GeneralResponse: &gigxRR.GeneralResponse{
			Message:            emailType,
			IsEmailSuccess:     true,
			IsOperationSuccess: true,
		},
	}, nil
}
