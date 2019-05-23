package service 

import (
	"time"
	"bytes"
	"fmt"
	"html/template"
	gomail "gopkg.in/gomail.v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
)
type UserRegisterData struct {
    WelcomeToGignax string
    ThanksForSigup string
    ClickVerrificationLink string
	VerificationLink string
	YourVerificationLink string
	OnceVerifiedAccount string
	UserVerifiedAccountInformation string
	OneUseLink string
	Email string
}
type UserForgotPasswordData struct {
    VerificationTokenTitle string
    ReceivedPasswordChangeRequest string
    ViaEmailAddress string
	DontShareVerificationLink string
	ResetPasswordLink string
	YourVerificationLink string 
	OneUseToken string
	Email string
}
func  SendUserRegisterConfirmationMailService(userEmail_ string,emailType_ string,verificationToken_ string,lang_ string) (*gigxRR.SendEmailResponse, error) {
	
	t := time.Now().UTC()
	userTemporaryInformation := repo.UserTemporaryInformation {
		Email: userEmail_,
		ForgotPasswordVerificationToken:"",
		RegisterVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		ForgotPasswordVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		EmailType:emailType_,
		IsTokenUsed: false,
    	IsTokenExpired: false, 
	}
	// if there is verification code not used
	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository=&userTemporaryInformation
	if dbResp,dbErr := userTemporaryInformationOp.CheckVerificationTokenResentEmail(); dbErr == nil {
		userTemporaryInformation.Id=dbResp.Id
		userTemporaryInformation.RegisterVerificationToken=dbResp.RegisterVerificationToken
		dbResp.IsTokenUsed=false
		dbResp.IsTokenExpired=true
		if dbErr := userTemporaryInformationOp.Update(); dbErr != nil {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(Translate(lang_,"resent_register_email_update_database_error")+":%v",dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.RegisterVerificationToken=verificationToken_
	mailTypePath:="app_root/mail_templates/user_register_confirmation.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_template_parse_error")+":%v",err.Error()),
		)
	}
	wd := UserRegisterData {
        WelcomeToGignax: Translate(lang_,"welcome_to_gignox"),
        ThanksForSigup: Translate(lang_,"thanks_for_signup"),
        ClickVerrificationLink: Translate(lang_,"click_verification_link"),
        YourVerificationLink: Translate(lang_,"your_verification_link"),
        OnceVerifiedAccount: Translate(lang_,"once_verified_user_account"),
        UserVerifiedAccountInformation: Translate(lang_,"verified_user_account_information"),
		VerificationLink: verificationToken_,
		Email:userEmail_,		
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_template_execute_error")+":%v",err.Error()),
		)
	}

	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"user_temporary_information_insert_error")+":%v",err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail_)
	mail.SetHeader("Subject", Translate(lang_,"user_register_mail_subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_send_error")+":%v",err.Error()),
		)
	}
	return &gigxRR.SendEmailResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:emailType_,
			IsEmailSuccess:true,
			IsOperationSuccess:true,
		},
	}, nil
}



func  SendUserForgotPasswordVerificationMailService(userEmail_ string,emailType_ string,verificationToken_ string,lang_ string) (*gigxRR.SendEmailResponse, error) {
	
	t := time.Now().UTC()
	userTemporaryInformation := repo.UserTemporaryInformation {
		Email: userEmail_,
		RegisterVerificationToken:"",
		RegisterVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		ForgotPasswordVerificationTokenCreateDate: t.Format("2006-01-02 15:04:05"),
		EmailType:emailType_,
		IsTokenUsed: false,
    	IsTokenExpired: false, 
	}
	// if there is verification code not used
	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository=&userTemporaryInformation
	if dbResp,dbErr := userTemporaryInformationOp.CheckVerificationTokenResentEmail(); dbErr == nil {
		userTemporaryInformation.Id=dbResp.Id
		userTemporaryInformation.ForgotPasswordVerificationToken=dbResp.ForgotPasswordVerificationToken
		userTemporaryInformation.IsTokenUsed=false
		userTemporaryInformation.IsTokenExpired=true
		if dbErr := userTemporaryInformationOp.Update(); dbErr != nil {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(Translate(lang_,"resent_forgot_password_email_update_database_error")+":%v",dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.ForgotPasswordVerificationToken=verificationToken_
	mailTypePath:="app_root/mail_templates/user_forgot_password.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_template_parse_error")+":%v",err.Error()),
		)
	}
	var resetPasswordLink = "http://localhost:3000/password_reset/" + verificationToken_;
	wd := UserForgotPasswordData{
        VerificationTokenTitle: Translate(lang_,"password_reset_token_title"),
        ReceivedPasswordChangeRequest: Translate(lang_,"received_password_change_request"),
        ViaEmailAddress: Translate(lang_,"via_email_address"),
		DontShareVerificationLink: Translate(lang_,"dont_share_verification_token"),
		YourVerificationLink: Translate(lang_,"your_password_reset_link"),
        OneUseToken: Translate(lang_,"one_use_link"),
		ResetPasswordLink: resetPasswordLink,
		Email:userEmail_,
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_template_execute_error")+":%v",err.Error()),
		)
	}
	userTemporaryInformation.IsTokenUsed=false
	userTemporaryInformation.IsTokenExpired=false
	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"user_temporary_information_insert_error")+":%v",err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail_)
	mail.SetHeader("Subject", Translate(lang_,"user_forgot_mail_subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"email_send_error")+":%v",err.Error()),
		)
	}
	return &gigxRR.SendEmailResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:emailType_,
			IsEmailSuccess:true,
			IsOperationSuccess:true,
		},
	}, nil
}
