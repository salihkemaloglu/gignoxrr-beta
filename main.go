package main

import (
	
	
	"fmt"
	"time"
	"context"
	"net/http"
	"encoding/hex"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"github.com/patrickmn/go-cache"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validation"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repository"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

type server struct {
}
type respCache struct {
	Count int 
}
var c *cache.Cache

func (s *server) SayHello(ctx context.Context, req *gigxRR.HelloRequest) (*gigxRR.HelloResponse, error) {
	
	fmt.Printf("RR service is working...Received rpc from client, message=%s\n", req.GetMessage())
	
	return &gigxRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
func (s *server) Login(ctx context.Context, req *gigxRR.LoginUserRequest) (*gigxRR.LoginUserResponse, error) {
	
	fmt.Printf("RR service is working for Login...Received rpc from client.\n")

	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	data := req.GetUser();
	user := db.User {
		Username:data.GetUsername(),
		Password:data.GetPassword(),
	}
	if res := val.UserLoginFieldValidation(user,lang); res != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(res),
		)
	}
	var loginAttemptCount int
	if x, found := c.Get(user.Username); found {
		loginAttemptCount = x.(int)
		if loginAttemptCount >= 7{
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"User_Login_attemps")),
			)
		}
	}
	var op repo.UserRepository=user
	if  err:= op.Login(); err != nil {
	    if x, found := c.Get(user.Username); found {
			loginAttemptCount = x.(int)
			if loginAttemptCount < 7 {
				loginAttemptCount=loginAttemptCount+1
				c.Set(user.Username, loginAttemptCount, cache.DefaultExpiration)
			}
	    } else {
			c.Set(user.Username, 1, cache.DefaultExpiration)
		}
		return nil,status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf(helper.Translate(lang,"Invalid_User_Information")),
		)
	}
	tokenRes,tokenErr:=helper.CreateTokenEndpoint(user)
	if tokenErr != nil{
		return nil,status.Errorf(
			codes.Unknown,
			fmt.Sprintf(helper.Translate(lang,"Token_Create_Error") +": %v",tokenErr.Error()),
		)
	}

	return &gigxRR.LoginUserResponse{
		User:&gigxRR.UserLogin{
			Username:	user.Username,
			Token:		tokenRes,
		},
	}, nil

}
func (s *server) Register(ctx context.Context, req *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {
	
	fmt.Printf("RR service is working for Register...Received rpc from client.\n")
	
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)
	verificationCode,verErr:=helper.GenerateVerificationCode()
	if verErr !=nil {
		verificationCode = "134584"
	}
	data := req.GetUser();
	t := time.Now().UTC()
	user := db.User {
		Name: data.GetName(),
		Surname: data.GetSurname(),
		Username: data.GetUsername(),
		Email: data.GetEmail(),
		Password: data.GetPassword(),
		CreatedDate: t.Format("2006-01-02 15:04:05"),
		UpdatedDate: t.Format("2006-01-02 15:04:05"),
		TotalSpace: 100,
		LanguageType: userLang,

		RegisterVerificationCode: verificationCode, 	   
		ForgotPasswordVerificationCode:"0" ,
	}
	
	if valResp := val.UserRegisterFieldValidation(user,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var op repo.UserRepository=user
	if err := op.CheckUser(); err ==nil  {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(helper.Translate(lang,"Already_Created_Account")+user.Username),
		)
	}

	if dbResp := op.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf(helper.Translate(lang,"Account_Insert_Error")+" :%v",dbResp.Error()),
		)
	}
	isOk:="ok"
	if mailResp:=helper.SendUserRegisterConfirmationMail(user.Email,userLang,verificationCode); mailResp != "ok" {
		isOk=mailResp
	}

	return &gigxRR.RegisterUserResponse{
		User:&gigxRR.User{
			Id:       	hex.EncodeToString([]byte(user.Id)),
			Username:	user.Username,
			IsMailSuccess:isOk,
		},
	}, nil
}
func (s *server) UpdateUser(ctx context.Context, req *gigxRR.UpdateUserRequest) (*gigxRR.UpdateUserResponse, error) {
	return nil,nil
}
func (s *server) DeleteUser(ctx context.Context, req *gigxRR.DeleteUserRequest) (*gigxRR.DeleteUserResponse, error) {
	return nil,nil
}
func (s *server) GetFile(ctx context.Context, req *gigxRR.GetFileRequest) (*gigxRR.GetFileResponse, error) {
	return nil,nil
}
func (s *server) GetAllFiles(req *gigxRR.GetAllFilesRequest, stream gigxRR.GigxRRService_GetAllFilesServer)error {
	return nil
}
func (s *server) UpdateFile(ctx context.Context, req *gigxRR.UpdateFileRequest) (*gigxRR.UpdateFileResponse, error) {
	return nil,nil
}
func (s *server) DeleteFile(ctx context.Context, req *gigxRR.DeleteFileRequest) (*gigxRR.DeleteFileResponse, error) {
	return nil,nil
}
var (

	// useWebsockets = pflag.Bool("use_websockets", false, "whether to use beta websocket transport layer")
	enableTls       = pflag.Bool("enable_tls", true, "Use TLS - required for HTTP2.")
	configFileType  = pflag.Bool("config_file_type", true, "true is for production,false is for local development")
	tlsCertFilePath = pflag.String("tls_cert_file", "app-root/ssl/fullchain.pem", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = pflag.String("tls_key_file", "app-root/ssl/privkey.pem", "Path to the private key file.")
	// flagHttpMaxWriteTimeout = pflag.Duration("server_http_max_write_timeout", 10*time.Second, "HTTP server config, max write duration.")
	// flagHttpMaxReadTimeout  = pflag.Duration("server_http_max_read_timeout", 10*time.Second, "HTTP server config, max read duration.")
)
func main(){
	pflag.Parse()

	port :=8900
	if *enableTls {
		port = 8901
	}

	fmt.Println("RR Service is Starting...")
	// init language folder path
	err := helper.InitLocales("app-root/languages")
	if err != nil {
		fmt.Println("Error happened when langs file loaded", err.Error())
	}
	// create new cache for user login attemtps
	c = cache.New(5*time.Minute, 10*time.Minute)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	gigxRR.RegisterGigxRRServiceServer(grpcServer, &server{})

	fmt.Println("Mongodb Service Started")
	if confErr:=db.LoadConfiguration(*configFileType); confErr!="ok"{
		fmt.Println(confErr)
	}

	allowedOrigins := helper.MakeAllowedOrigins()
	
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(helper.MakeHttpOriginFunc(allowedOrigins)),
	}

	wrappedGrpc := grpcweb.WrapServer(grpcServer, options...)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server. http port: %d, with TLS: %v", port, *enableTls)


	if *enableTls {
		fmt.Printf("server started as  https and listen to port: %v \n",port)
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		fmt.Printf("server started as http and listen to port: %v \n",port)
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}
}
