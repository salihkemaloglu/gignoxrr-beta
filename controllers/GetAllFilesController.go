package controllers

import (
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
)

//GetAllFiles ...
func (s *Server) GetAllFiles(req *gigxRR.GetAllFilesRequest, stream gigxRR.GigxRRService_GetAllFilesServer) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return nil
}
