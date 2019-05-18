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
    MustEnterVerrificationToken string
	VerificationToken string
	YourVerificationToken string
	OneUseToken string
}
type UserForgotPasswordData struct {
    VerificationTokenTitle string
    ReceivedPasswordChangeRequest string
    ViaEmailAddress string
	DontShareVerificationToken string
	ResetPasswordLink string
	YourVerificationToken string 
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
				fmt.Sprintf(Translate(lang_,"Resent_Register_Email_Update_Database_Error")+":%v",dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.RegisterVerificationToken=verificationToken_
	mailTypePath:="app_root/mail_templates/user_register_confirmation.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Template_Parse_Error")+":%v",err.Error()),
		)
	}
	wd := UserRegisterData{
        WelcomeToGignax: Translate(lang_,"Welcome_To_Gignox"),
        ThanksForSigup: Translate(lang_,"Thanks_For_Sigup"),
        MustEnterVerrificationToken: Translate(lang_,"Must_Enter_Verrification_Token"),
        YourVerificationToken: Translate(lang_,"Your_Verification_Url"),
        OneUseToken: Translate(lang_,"One_Use_Token"),
        VerificationToken: verificationToken_,
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Template_Execute_Error")+":%v",err.Error()),
		)
	}

	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"User_Temporary_Indormation_Insert_Error")+":%v",err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail_)
	mail.SetHeader("Subject", Translate(lang_,"User_Register_Mail_Subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Send_Error")+":%v",err.Error()),
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
				fmt.Sprintf(Translate(lang_,"Resent_Forgot_Password_Email_Update_Database_Error")+":%v",dbErr.Error()),
			)
		}

	}
	userTemporaryInformation.ForgotPasswordVerificationToken=verificationToken_
	mailTypePath:="app_root/mail_templates/user_forgot_password.html"
	temp, err := template.ParseFiles(mailTypePath)
	if err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Template_Parse_Error")+":%v",err.Error()),
		)
	}
	var resetPasswordLink = "http://localhost:3000/password_reset/" + verificationToken_;
	wd := UserForgotPasswordData{
        VerificationTokenTitle: Translate(lang_,"Password_reset_Token_Title"),
        ReceivedPasswordChangeRequest: Translate(lang_,"Received_Password_Change_Request"),
        ViaEmailAddress: Translate(lang_,"Via_Email_Address"),
		DontShareVerificationToken: Translate(lang_,"Dont_Share_Verification_Token"),
		YourVerificationToken: Translate(lang_,"Your_Password_Reset_Url"),
        OneUseToken: Translate(lang_,"One_Use_Token"),
		ResetPasswordLink: resetPasswordLink,
		Email:userEmail_,
    }

	var mailBytes bytes.Buffer
	if err := temp.Execute(&mailBytes, &wd); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Template_Execute_Error")+":%v",err.Error()),
		)
	}
	userTemporaryInformation.IsTokenUsed=false
	userTemporaryInformation.IsTokenExpired=false
	if dbResp := userTemporaryInformationOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"User_Temporary_Indormation_Insert_Error")+":%v",err.Error()),
		)
	}

	result := mailBytes.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", "gignox.us@gmail.com")
	mail.SetHeader("To", userEmail_)
	mail.SetHeader("Subject", Translate(lang_,"User_Forgot_Mail_Subject"))//bu da dinamik olacak
	mail.SetBody("text/html", result)
	// mail.Attach(mailTypePath)
	
	dial := gomail.NewDialer("smtp.gmail.com", 587, "gignox.us@gmail.com", "mameguli13--")

	// Send the email to user
	if err := dial.DialAndSend(mail); err != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(Translate(lang_,"Email_Send_Error")+":%v",err.Error()),
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
