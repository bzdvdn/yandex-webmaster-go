package yandexwebmaster

import (
	"fmt"
	"time"
)

// Indexing service for indexing management
type IndexingService struct {
	client *Client
}

// func for init IndexingService
func newIndexingService(cl *Client) *IndexingService {
	return &IndexingService{client: cl}
}

type Indicator struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

type Indicators struct {
	Indicators map[string]*Indicator `json:"inidicators"`
}

type Sample struct {
	Status     string `json:"status"`
	HTTPCode   string `json:"http_code"`
	URL        string `json:"url"`
	AccessDate string `json:"access_date"`
}

type SamplesResult struct {
	Count   int       `json:"count"`
	Samples []*Sample `json:"samples"`
}

// get indexing history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-indexing-history.html
func (s *IndexingService) GetIndexingHistory(hostID string, dateFrom time.Time, dateTo time.Time) (Indicators, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	var result Indicators
	endpoint := fmt.Sprintf("user/%d/hosts/%s/indexing/history", s.client.userID, hostID)
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get indexing samples, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/hosts-indexing-samples.html
func (s *IndexingService) GetIndexingSamples(hostID string, limit int, offset int) (SamplesResult, error) {
	data := make(map[string]interface{})
	data["limit"] = limit
	data["offset"] = offset

	var result SamplesResult
	endpoint := fmt.Sprintf("user/%d/hosts/%s/indexing/samples", s.client.userID, hostID)
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}
