package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
)

//UploadFileService ...
func UploadFileService(stream gigxRR.GigxRRService_UploadFileServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		file := req.GetFile()
		// copy example
		absPath, _ := filepath.Abs("handlersss.mp4")
		f, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)

		}
		defer f.Close() //if there is a bug. If the call to os.Create fails, the function will return without closing the source file. defer is closing it.
		reader := bytes.NewReader(file)
		io.Copy(f, reader)

		sendErr := stream.Send(&gigxRR.UploadFileResponse{
			Result: "ok",
		})
		if sendErr != nil {
			fmt.Println("Error while sending data to client:", sendErr)
			return sendErr
		}
		fmt.Printf("LongGreet function was close with a streaming request\n")
	}
}
