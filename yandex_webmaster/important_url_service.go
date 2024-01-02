package yandexwebmaster

import (
	"fmt"
	"net/http"
)

// Important url service for important url management
type ImportantURLService struct {
	client *Client
}

// func for init ImportantURLService
func newImportURLService(cl *Client) *ImportantURLService {
	return &ImportantURLService{client: cl}
}

type IndexingStatus struct {
	Status     string `json:"status"`
	HTTPCode   string `json:"http_code"`
	AccessDate string `json:"access_date"`
}

type SearchStatus struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	LastAccess        string `json:"last_access"`
	ExcludedURLStatus string `json:"excluded_url_status"`
	BadHTTPStatus     int    `json:"bad_http_status"`
	Searchable        bool   `json:"searchable"`
	TargetURL         string `json:"target_url"`
}

type ImportantURL struct {
	URL              string          `json:"url"`
	UpdateDate       string          `json:"update_date"`
	ChangeIndicators []string        `json:"change_indicators"`
	IndexingStatus   *IndexingStatus `json:"indexing_status"`
	SearchStatus     *SearchStatus   `json:"search_status"`
}

type ImportantURLS struct {
	URLS []*ImportantURL `json:"urls"`
}

type ImportantURLSHistory struct {
	History []*ImportantURL `json:"history"`
}

// get monigorint important urls, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-id-important-urls.html
func (s *IndexingService) GetMonitoringImportantURLS(hostID string) (ImportantURLS, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/important-urls", s.client.userID, hostID)
	var result ImportantURLS
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

// get important url history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-id-important-urls-history.html
func (s *IndexingService) GetImportantURLHistory(hostID string, url string) (ImportantURLSHistory, error) {
	data := make(map[string]interface{})
	data["url"] = url
	endpoint := fmt.Sprintf("user/%d/hosts/%s/important-urls", s.client.userID, hostID)
	var result ImportantURLSHistory
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}
