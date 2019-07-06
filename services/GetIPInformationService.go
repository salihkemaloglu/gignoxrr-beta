package services

import (
	"context"

	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
)

//GetIPInformationService ...
func GetIPInformationService(ctx context.Context) (*gigxRR.GetIPInformationResponse, error) {

	return helper.GetIPInformation(ctx, true)
}
