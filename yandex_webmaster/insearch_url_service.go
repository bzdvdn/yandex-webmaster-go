package yandexwebmaster

import (
	"fmt"
	"time"
)

// Insearch url service for insearch url management
type InsearchURLService struct {
	client *Client
}

// func for init InsearchURLService
func newInsearchURLService(cl *Client) *InsearchURLService {
	return &InsearchURLService{client: cl}
}

type InsearchURLHistoryData struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

type InseacrhURLHistory struct {
	History []*InsearchURLHistoryData `json:"history"`
}

type InsearchSample struct {
	URL        string `json:"url"`
	LastAccess string `json:"last_access"`
	Title      string `json:"title"`
}

type InsearchSampleResponse struct {
	Count   int               `json:"count"`
	Samples []*InsearchSample `json:"samples"`
}

type SearchURLEventHistory struct {
	AppeadINSearch    []*Indicator `json:"APPEARED_IN_SEARCH"`
	RemovedFromSearch []*Indicator `json:"REMOVED_FROM_SEARCH"`
}

type SearchURLEventHistoryResponse struct {
	Indicators SearchURLEventHistory `json:"indicators"`
}

type InsearchEventSample struct {
	URL               string `json:"url"`
	Title             string `json:"title"`
	EventDate         string `json:"event_date"`
	LastAccess        string `json:"last_access"`
	Event             string `json:"event"`
	ExcludedURLStatus string `json:"excluded_url_status"`
	BadHTTPStatus     int    `json:"bad_http_status"`
	TargetURL         string `json:"target_url"`
}

type InsearchEventSampleResponse struct {
	Samples []*InsearchEventSample `json:"InsearchEventSample"`
}

// get insearch url history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-indexing-insearch-history.html
func (s *InsearchURLService) GetInsearchURLHistory(hostID string, dateFrom time.Time, dateTo time.Time) (InseacrhURLHistory, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-urls/in-search/history", s.client.userID, hostID)
	var result InseacrhURLHistory
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get insearch url samples, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-indexing-insearch-samples.html
func (s *InsearchURLService) GetInsearchURLSamples(hostID string, limit int, offset int) (InsearchSampleResponse, error) {
	data := make(map[string]interface{})
	data["limit"] = limit
	data["offset"] = offset
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-urls/in-search/samples", s.client.userID, hostID)
	var result InsearchSampleResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get insearch url events history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-search-events-history.html
func (s *InsearchURLService) GetInsearchURLEventsHistory(hostID string, dateFrom time.Time, dateTo time.Time) (SearchURLEventHistoryResponse, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-urls/events/history", s.client.userID, hostID)
	var result SearchURLEventHistoryResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get insearch url event samples, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-search-events-samples.html
func (s *InsearchURLService) GetInsearchURLEventSamples(hostID string, limit int, offset int) (InsearchEventSampleResponse, error) {
	data := make(map[string]interface{})
	data["limit"] = limit
	data["offset"] = offset
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-urls/events/samples", s.client.userID, hostID)
	var result InsearchEventSampleResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}
