package controller

import (
	"context"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//GetIPInformationController ...
func GetIPInformationController(ctx context.Context) (*gigxRR.GetIPInformationResponse, error) {

	return helper.GetIPInformation(ctx, true)
}
