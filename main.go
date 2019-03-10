package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/salihkemaloglu/DemRR-beta-001/proto"
	"github.com/salihkemaloglu/DemRR-beta-001/middleware"
	"github.com/salihkemaloglu/DemRR-beta-001/proxy"
	repo "github.com/salihkemaloglu/DemRR-beta-001/repository"
	db "github.com/salihkemaloglu/DemRR-beta-001/mongodb"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *demRR.HelloRequest) (*demRR.HelloResponse, error) {
	fmt.Printf("RR service is working...Received rpc from client, message=%s\n", req.GetMessage())

	data :=db.File{Name:req.GetMessage()}
	var op repo.FileRepository =data
	var items, _ = op.Insert()
	fmt.Println("Received a message:", items)

	return &demRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
func (s *server) Login(ctx context.Context, req *demRR.LoginUserRequest) (*demRR.LoginUserResponse, error) {
	return nil,nil
}
func (s *server) Register(ctx context.Context, req *demRR.RegisterUserRequest) (*demRR.RegisterUserResponse, error) {
	db.UserInformation
	data :=req.GetUser();
	info:=req.GetUserInformation();
	// var op repo.FileRepository =data
	// var items, _ = op.Insert()
	fmt.Println("username:", data.GetUsername())
	fmt.Println("password:", data.GetPassword())
	fmt.Println("info:", info.GetDescription())
	return &demRR.RegisterUserResponse{
		User:&demRR.User{
			Id:       	"oid.Hex()",
			Username:	data.GetUsername(),
			Password:	data.GetPassword(), 	
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
func (s *server) GetAllFiles(req *demRR.GetAllFilesRequest, stream demRR.DemService_GetAllFilesServer)error {
	data := &blogItem{}
	stream.Send(&demRR.GetAllFilesResponse{File: dataToBlogPb(data)})
	return nil
}
func (s *server) UpdateFile(ctx context.Context, req *demRR.UpdateFileRequest) (*demRR.UpdateFileResponse, error) {
	return nil,nil
}
func (s *server) DeleteFile(ctx context.Context, req *demRR.DeleteFileRequest) (*demRR.DeleteFileResponse, error) {
	return nil,nil
}
type blogItem struct {
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}
func dataToBlogPb(data *blogItem) *demRR.File {
	return &demRR.File{
		Name: data.AuthorID,
	}
}
func main(){

	fmt.Println("RR Service Started")
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	demRR.RegisterDemServiceServer(grpcServer, &server{})
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