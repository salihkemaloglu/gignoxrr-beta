package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/patrickmn/go-cache"
	ctrl "github.com/salihkemaloglu/gignoxrr-beta-001/controllers"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var c *cache.Cache

func main() {
	pflag.Parse()

	fmt.Println("RR Service is Starting...")
	// init languagecode folder path
	err := helper.InitLocales("localization/languages")
	if err != nil {
		fmt.Println("Error happened when langs file loaded", err.Error())
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	// create new cache for user login attemtps
	c = cache.New(5*time.Minute, 10*time.Minute)

	opts := []grpc.ServerOption{}
	// opts = append(opts, grpc.MaxRecvMsgSize(1024*1024*1024))
	grpcServer := grpc.NewServer(opts...)
	gigxRR.RegisterGigxRRServiceServer(grpcServer, &ctrl.Server{})

	fmt.Println("Mongodb Service Started")
	repo.LoadConfiguration("dev")

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
		Addr:    fmt.Sprintf(":%v", 8901),
		Handler: http.HandlerFunc(handler),
	}

	fmt.Printf("server started as http and listen to port: %v \n", 8901)
	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}

}
