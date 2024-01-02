package yandexwebmaster

import (
	"fmt"
	"net/http"
)

// Host service for host management
type HostService struct {
	client *Client
}

// func for init new HostService
func newHostService(cl *Client) *HostService {
	return &HostService{client: cl}
}

// MainMirror strunct
type MainMirror struct {
	HostID         string `json:"host_id"`
	AsciiHostURL   string `json:"ascii_host_url"`
	UnicodeHostURL string `json:"unicode_host_url"`
	Verified       bool   `json:"verified"`
}

// Host Struct
type Host struct {
	HostID         string     `json:"host_id"`
	AsciiHostURL   string     `json:"ascii_host_url"`
	UnicodeHostURL string     `json:"unicode_host_url"`
	Verified       bool       `json:"verified"`
	MainMirror     MainMirror `json:"main_mirror"`
}

// hosts response struct
type Hosts struct {
	Hosts []*Host `json:"hosts"`
}

type CreatedHost struct {
	HostURL string `json:"host_url"`
}

// get hosts from yandex webmaster, DOC: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts.html
func (s *HostService) GetHosts() (Hosts, error) {
	endpoint := fmt.Sprintf("user/%d/hosts", s.client.userID)
	var result Hosts
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

func (s *HostService) GetHost(hostID string) (Host, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s", s.client.userID, hostID)
	var result Host
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

func (s *HostService) AddHost(hostURL string) (CreatedHost, error) {
	endpoint := fmt.Sprintf("user/%d/hosts", s.client.userID)
	data := make(map[string]interface{})
	data["host_url"] = hostURL
	var result CreatedHost
	_, err := s.client.sendAPIRequest(http.MethodPost, endpoint, data, &result)
	return result, err

}

func (s *HostService) DeleteHost(hostID string) (interface{}, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s", s.client.userID, hostID)
	var result interface{}
	_, err := s.client.sendAPIRequest(http.MethodDelete, endpoint, nil, &result)
	return result, err
}
