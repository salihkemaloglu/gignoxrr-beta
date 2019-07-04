package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/patrickmn/go-cache"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type server struct {
}

var (
	c *cache.Cache
	// useWebsockets = pflag.Bool("use_websockets", false, "whether to use beta websocket transport layer")
	enableTLS       = pflag.Bool("enable_tls", false, "Use TLS - required for HTTP2.") // false is for local development
	tlsCertFilePath = pflag.String("tls_cert_file", "app_root/ssl/fullchain.pem", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = pflag.String("tls_key_file", "app_root/ssl/privkey.pem", "Path to the private key file.")
	// flagHttpMaxWriteTimeout = pflag.Duration("server_http_max_write_timeout", 10*time.Second, "HTTP server config, max write duration.")
	// flagHttpMaxReadTimeout  = pflag.Duration("server_http_max_read_timeout", 10*time.Second, "HTTP server config, max read duration.")
)

func (s *server) SayHello(ctx context.Context, req *gigxRR.HelloRequest) (*gigxRR.HelloResponse, error) {

	fmt.Printf("RR service is working for SayHello...Received rpc from client, message=%s\n", req.GetMessage())
	return &gigxRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
func (s *server) GetIPInformation(ctx context.Context, req *gigxRR.GetIPInformationRequest) (*gigxRR.GetIPInformationResponse, error) {

	fmt.Printf("RR service is working for GetIpAddess...Received rpc from client")
	return serv.GetIPInformationService(ctx)
}
func (s *server) Login(ctx context.Context, req *gigxRR.LoginUserRequest) (*gigxRR.LoginUserResponse, error) {

	fmt.Printf("RR service is working for Login...Received rpc from client.\n")
	return serv.LoginService(ctx, req, c)
}
func (s *server) Register(ctx context.Context, req *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {

	fmt.Printf("RR service is working for Register...Received rpc from client.\n")
	return serv.RegisterService(ctx, req)
}
func (s *server) CheckUserToRegister(ctx context.Context, req *gigxRR.CheckUserToRegisterRequest) (*gigxRR.CheckUserToRegisterResponse, error) {

	fmt.Printf("RR service is working for CheckUserToRegister...Received rpc from client.\n")
	return serv.CheckUserToRegisterService(ctx, req)
}
func (s *server) SendEmail(ctx context.Context, req *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {

	fmt.Printf("RR service is working for SendMail...Received rpc from client.\n")
	return serv.SendEmailService(ctx, req)
}
func (s *server) CheckVerificationLink(ctx context.Context, req *gigxRR.CheckVerificationLinkRequest) (*gigxRR.CheckVerificationLinkResponse, error) {

	fmt.Printf("RR service is working for CheckVerificationLink...Received rpc from client.\n")
	return serv.CheckVerificationLinkService(ctx, req)
}
func (s *server) ResetUserPassword(ctx context.Context, req *gigxRR.ResetUserPasswordRequest) (*gigxRR.ResetUserPasswordResponse, error) {

	fmt.Printf("RR service is working for ResetUserPassword...Received rpc from client.\n")
	return serv.ResetUserPasswordService(ctx, req)
}
func (s *server) GetUser(ctx context.Context, req *gigxRR.GetUserRequest) (*gigxRR.GetUserResponse, error) {

	fmt.Printf("RR service is working for GetUser...Received rpc from client.\n")
	return serv.GetUserService(ctx, req)
}
func (s *server) UpdateUser(ctx context.Context, req *gigxRR.UpdateUserRequest) (*gigxRR.UpdateUserResponse, error) {
	return nil, nil
}
func (s *server) DeleteUser(ctx context.Context, req *gigxRR.DeleteUserRequest) (*gigxRR.DeleteUserResponse, error) {
	return nil, nil
}
func (s *server) GetFile(ctx context.Context, req *gigxRR.GetFileRequest) (*gigxRR.GetFileResponse, error) {
	return nil, nil
}
func (s *server) GetAllFiles(req *gigxRR.GetAllFilesRequest, stream gigxRR.GigxRRService_GetAllFilesServer) error {
	return nil
}
func (s *server) UpdateFile(ctx context.Context, req *gigxRR.UpdateFileRequest) (*gigxRR.UpdateFileResponse, error) {
	return nil, nil
}
func (s *server) DeleteFile(ctx context.Context, req *gigxRR.DeleteFileRequest) (*gigxRR.DeleteFileResponse, error) {
	return nil, nil
}
func (s *server) UploadFile(stream gigxRR.GigxRRService_UploadFileServer) error {
	fmt.Printf("RR service is working for UploadFile...Received rpc from client.\n")
	return serv.UploadFileService(stream)
}

func main() {
	pflag.Parse()

	port := 8900
	configFile := "dev"
	if *enableTLS {
		configFile = "prod"
		port = 8901
	}

	fmt.Println("RR Service is Starting...")
	// init languagecode folder path
	err := helper.InitLocales("app_root/languages")
	if err != nil {
		fmt.Println("Error happened when langs file loaded", err.Error())
	}
	// create new cache for user login attemtps
	c = cache.New(5*time.Minute, 10*time.Minute)

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.MaxRecvMsgSize(1024*1024*1024))
	grpcServer := grpc.NewServer(opts...)
	gigxRR.RegisterGigxRRServiceServer(grpcServer, &server{})

	fmt.Println("Mongodb Service Started")
	if confErr := repo.LoadConfiguration(configFile); confErr != "ok" {
		fmt.Println(confErr)
	}

	allowedOrigins := helper.MakeAllowedOrigins()

	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(helper.MakeHTTPOriginFunc(allowedOrigins)),
	}

	wrappedGrpc := grpcweb.WrapServer(grpcServer, options...)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server. http port: %d, with TLS: %v", port, *enableTLS)

	if *enableTLS {
		fmt.Printf("server started as  https and listen to port: %v \n", port)
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		fmt.Printf("server started as http and listen to port: %v \n", port)
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}
}
