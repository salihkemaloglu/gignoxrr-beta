package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// IPStacInformation ...
type IPStacInformation struct {
	IP            string  `json:"ip"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Location      struct {
		GeonameID int    `json:"geoname_id"`
		Capital   string `json:"capital"`
		Languages []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		CallingCode             string `json:"calling_code"`
		IsEu                    bool   `json:"is_eu"`
	} `json:"location"`
}

// GeonameInformation ...
type GeonameInformation struct {
	Sunrise     string  `json:"sunrise"`
	Lng         float64 `json:"lng"`
	CountryCode string  `json:"countryCode"`
	GmtOffset   int32   `json:"gmtOffset"`
	RawOffset   int     `json:"rawOffset"`
	Sunset      string  `json:"sunset"`
	TimezoneID  string  `json:"timezoneId"`
	DstOffset   int     `json:"dstOffset"`
	CountryName string  `json:"countryName"`
	Time        string  `json:"time"`
	Lat         float64 `json:"lat"`
}

//GetIPInformation ...
func GetIPInformation(ctx context.Context, requestType bool) (*gigxRR.GetIPInformationResponse, error) {

	pr, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf("getClinetIP, invoke FromContext() failed"),
		)
	}
	if pr.Addr == net.Addr(nil) {
		return nil, status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf("getClientIP, peer.Addr is nil"),
		)
	}

	subIPAddress := "85.108.130.101"
	if len(pr.Addr.String()) > 15 {
		subIPAddress = pr.Addr.String()[0 : len(pr.Addr.String())-6]
	}

	if requestType {
		ipstack, ipstackError := getIPInformationFromIpstack(subIPAddress)
		if ipstackError != "ok" {
			return nil, status.Errorf(
				codes.Unimplemented,
				fmt.Sprintf(ipstackError),
			)
		}
		geoname, geonameError := getLocationFromGeoname(ipstack.Latitude, ipstack.Longitude)
		if geonameError != "ok" {
			return nil, status.Errorf(
				codes.Unimplemented,
				fmt.Sprintf(geonameError),
			)
		}
		return &gigxRR.GetIPInformationResponse{
			IpInformation: &gigxRR.IPInformation{
				IpAddress:    subIPAddress,
				LanguageCode: ipstack.Location.Languages[0].Code,
				CountryFlag:  ipstack.Location.CountryFlag,
				CountryCode:  ipstack.CountryCode,
				CountryName:  ipstack.CountryName,
				GmtOffSet:    geoname.GmtOffset,
			},
		}, nil
	}
	return &gigxRR.GetIPInformationResponse{
		IpInformation: &gigxRR.IPInformation{
			IpAddress: subIPAddress,
		},
	}, nil

}

//getIPInformationFromIpstack ...
func getIPInformationFromIpstack(ipAddress string) (*IPStacInformation, string) {

	var url = "http://api.ipstack.com/" + ipAddress + "?access_key=373987625c97d99d00d294ea56da82d4"
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Sprintf("ipstack ip request error: %v", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Sprintf("ipstack response body error: %v", err.Error())
	}
	var info IPStacInformation
	errMarshal := json.Unmarshal(contents, &info)
	if errMarshal != nil {
		return nil, fmt.Sprintf("ipstack Unmarshal error: %v", errMarshal.Error())
	}
	return &info, "ok"

}

//getLocationFromGeoname ...
func getLocationFromGeoname(latitude float64, longitude float64) (*GeonameInformation, string) {

	var url = "http://api.geonames.org/timezoneJSON?lat=" + fmt.Sprintf("%f", latitude) + "&lng=" + fmt.Sprintf("%f", longitude) + "&username=gignox"
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Sprintf("geoname ip request error: %v", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Sprintf("ipstack response body error: %v", err.Error())
	}
	var info GeonameInformation
	errMarshal := json.Unmarshal(contents, &info)
	if errMarshal != nil {
		return nil, fmt.Sprintf("ipstack Unmarshal error: %v", errMarshal.Error())
	}
	return &info, "ok"

}
