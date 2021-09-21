package sms

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	Mock  = "mock"
	Debug = "debugging proxy"
	Prod  = "production"
)

type Service struct {
	APIKey string
	Env    string
}

func NewService(apiKey, env string) Service {
	return Service{apiKey, env}
}

// get url
func getUrl(env string) string {
	switch env {
	case Debug:
		return "https://private-f51dc7-unifysmsapi2.apiary-proxy.com/api/v2/sms/public/sendMessage"
	case Mock:
		return "https://private-f51dc7-unifysmsapi2.apiary-mock.com/api/v2/sms/public/sendMessage"
	case Prod:
		return "https://unify-base.rancard.com/api/v2/sms/public/sendMessage"
	default:
		return "https://unify-base.rancard.com/api/v2/sms/public/sendMessage"
	}

}

// SendGet
// Send - GET
func (service Service) SendGet(request SmsRequest1) (*SmsResponse, error) {
	client := &http.Client{}

	request.ApiKey = service.APIKey

	// map the actual post http request
	url := fmt.Sprintf(
		"%s?message=%s&msisdn=%s&apiKey=%s&scheduled=%v",
		getUrl(service.Env), request.Message, request.Msisdn, request.ApiKey, request.Scheduled,
	)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	// defer
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode == 200 {
		smsResponse, err := UnmarshalSmsResponse(res.Body)
		if err != nil {
			return nil, err
		}
		return &smsResponse, nil
	} else if strings.Contains(res.Status, "40") {
		// 40... status code
		return nil, errors.New("check your api key")
	} else {
		return nil, errors.New("unknown error occurred")
	}

}

// SendPost
// Send - POST
func (service Service) SendPost(request SmsRequest) (*SmsResponse, error) {
	client := &http.Client{}

	request.ApiKey = service.APIKey
	requestJSON, err := request.Marshal()
	if err != nil {
		return nil, err
	}
	// map the actual post http request
	res, err := client.Post(getUrl(service.Env), "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}
	// defer
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode == 200 {
		smsResponse, err := UnmarshalSmsResponse(res.Body)
		if err != nil {
			return nil, err
		}
		return &smsResponse, nil
	} else if strings.Contains(res.Status, "40") {
		// 40... status code
		return nil, errors.New("check your api key")
	} else {
		return nil, errors.New("unknown error occurred")
	}

}
