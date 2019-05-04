package controller

import (
	
	"context"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)
func GetIpInformationController(ctx_ context.Context) (*gigxRR.GetIpInformationResponse, error) {
	
   return helper.GetIpInformation(ctx_,true)
}
