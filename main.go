package main

import (
	
	"fmt"
	"context"
	"net/http"
	"encoding/hex"
	"github.com/rs/cors"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/grpclog"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/salihkemaloglu/DemRR-beta-001/proto"
	"github.com/salihkemaloglu/DemRR-beta-001/proxy"
	"github.com/salihkemaloglu/DemRR-beta-001/middleware"
	db "github.com/salihkemaloglu/DemRR-beta-001/mongodb"
	val "github.com/salihkemaloglu/DemRR-beta-001/validation"
	repo "github.com/salihkemaloglu/DemRR-beta-001/repository"
	token "github.com/salihkemaloglu/DemRR-beta-001/token"
	
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *demRR.HelloRequest) (*demRR.HelloResponse, error) {
	fmt.Printf("RR service is working...Received rpc from client, message=%s\n", req.GetMessage())
	return &demRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
func (s *server) Login(ctx context.Context, req *demRR.LoginUserRequest) (*demRR.LoginUserResponse, error) {
	fmt.Printf("RR service is working for Login...Received rpc from client.\n")
	data := req.GetUser();
	user := db.User {
		Username:data.GetUsername(),
		Password:data.GetPassword(),
	}
	if res := val.UserLoginFieldValidation(user); res != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(res),
		)
	}
	var op repo.UserRepository=user
	if err := op.Login(); err != nil {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Invalid Username or Password"),
		)
	}
	tokenRes,tokenErr:=token.CreateTokenEndpoint(user)
	if tokenErr != nil{
		return nil,status.Errorf(
			codes.Unknown,
			fmt.Sprintf("Token could not be created: %v",tokenErr.Error()),
		)
	}

	return &demRR.LoginUserResponse{
		User:&demRR.UserLogin{
			Username:	user.Username,
			Token:		tokenRes,
		},
	}, nil

}
func (s *server) Register(ctx context.Context, req *demRR.RegisterUserRequest) (*demRR.RegisterUserResponse, error) {
	
	fmt.Printf("RR service is working for Register...Received rpc from client.\n")
	
	data := req.GetUser();
	user := db.User {
		Username:data.GetUsername(),
		Password:data.GetPassword(),
		Email:data.GetEmail(),
	}

	if res := val.UserRegisterFieldValidation(user); res != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(res),
		)
	}

	var op repo.UserRepository=user
	if err := op.CheckUser(); err ==nil  {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("There is alreasy an account for: "+user.Username),
		)
	}

	resp := op.Insert()
	if resp != nil{
		return nil,status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf("Error happened when insert user :%v",resp.Error()),
		)
	}

	return &demRR.RegisterUserResponse{
		User:&demRR.User{
			Id:       	hex.EncodeToString([]byte(user.Id)),
			Username:	user.Username,
		},
	}, nil
}
func (s *server) UpdateUser(ctx context.Context, req *demRR.UpdateUserRequest) (*demRR.UpdateUserResponse, error) {
	return nil,nil
}
func (s *server) DeleteUser(ctx context.Context, req *demRR.DeleteUserRequest) (*demRR.DeleteUserResponse, error) {
	return nil,nil
}
func (s *server) GetFile(ctx context.Context, req *demRR.GetFileRequest) (*demRR.GetFileResponse, error) {
	return nil,nil
}
func (s *server) GetAllFiles(req *demRR.GetAllFilesRequest, stream demRR.DemRRService_GetAllFilesServer)error {
	return nil
}
func (s *server) UpdateFile(ctx context.Context, req *demRR.UpdateFileRequest) (*demRR.UpdateFileResponse, error) {
	return nil,nil
}
func (s *server) DeleteFile(ctx context.Context, req *demRR.DeleteFileRequest) (*demRR.DeleteFileResponse, error) {
	return nil,nil
}

func main(){

	fmt.Println("RR Service Started")
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	demRR.RegisterDemRRServiceServer(grpcServer, &server{})
	fmt.Println("Mongodb Service Started")
	db.LoadConfiguration()
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		middleware. NewGrpcWebMiddleware(wrappedGrpc).Handler,// Must come before general CORS handling
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}).Handler,
	)

	router.Get("/article-proxy", proxy.Article)

	if err := http.ListenAndServe(":8902", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}