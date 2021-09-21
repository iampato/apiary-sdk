package sms

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type SmsRequest struct {
	ApiKey          string   `json:"apiKey"`
	Contacts        []string `json:"contacts"`
	Message         string   `json:"message"`
	Scheduled       bool     `json:"scheduled"`
	HasPlaceholders bool     `json:"hasPlaceholders,omitempty"`
	SenderId        string   `json:"senderId"`
	StartDate       string   `json:"startDate,omitempty"`
	EndDate         string   `json:"endDate,omitempty"`
	GroupId         string   `json:"groupId,omitempty"`
}
type SmsRequest1 struct {
	ApiKey          string `json:"apiKey"`
	Msisdn          string `json:"msisdn"`
	Message         string `json:"message"`
	Scheduled       bool   `json:"scheduled"`
	HasPlaceholders bool   `json:"hasPlaceholders,omitempty"`
	SenderId        string `json:"senderId"`
	StartDate       string `json:"startDate,omitempty"`
	EndDate         string `json:"endDate,omitempty"`
	GroupId         string `json:"groupId,omitempty"`
}

func (r SmsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SmsResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func UnmarshalSmsResponse(reader io.Reader) (SmsResponse, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return SmsResponse{}, err
	}
	var r SmsResponse
	err = json.Unmarshal(data, &r)
	return r, err
}
