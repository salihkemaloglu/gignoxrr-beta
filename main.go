package main

import (
	"fmt"
	"time"
	"context"
	"net/http"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc/metadata"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
	cont "github.com/salihkemaloglu/gignox-rr-beta-001/controllers"
)

type server struct {
}
var c *cache.Cache

func (s *server) SayHello(ctx context.Context, req *gigxRR.HelloRequest) (*gigxRR.HelloResponse, error) {
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println(headers["user-agent"][0])
	}
	fmt.Printf("RR service is working...Received rpc from client, message=%s\n", req.GetMessage())
	return &gigxRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
func (s *server) Login(ctx context.Context, req *gigxRR.LoginUserRequest) (*gigxRR.LoginUserResponse, error) {
	
	fmt.Printf("RR service is working for Login...Received rpc from client.\n")
	return cont.LoginController(ctx,req,c)
}
func (s *server) Register(ctx context.Context, req *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {
	
	fmt.Printf("RR service is working for Register...Received rpc from client.\n")
	return cont.RegisterController(ctx,req)	
}
func (s *server) SendEmail(ctx context.Context, req *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {
	
	fmt.Printf("RR service is working for SendMail...Received rpc from client.\n")
	return cont.SendEmailController(ctx,req)		
}
func (s *server) CheckVerificationCode(ctx context.Context, req *gigxRR.CheckVerificationCodeRequest) (*gigxRR.CheckVerificationCodeResponse, error) {

	fmt.Printf("RR service is working for CheckVerificationCode...Received rpc from client.\n")
	return cont.CheckVerificationCodeController(ctx,req)	
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
	enableTls       = pflag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
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
	err := helper.InitLocales("app_root/languages")
	if err != nil {
		fmt.Println("Error happened when langs file loaded", err.Error())
	}
	// create new cache for user login attemtps
	c = cache.New(5*time.Minute, 10*time.Minute)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	gigxRR.RegisterGigxRRServiceServer(grpcServer, &server{})

	fmt.Println("Mongodb Service Started")
	if confErr:=db.LoadConfiguration("dev"); confErr!="ok"{
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
