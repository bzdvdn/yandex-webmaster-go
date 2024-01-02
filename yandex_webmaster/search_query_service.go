package yandexwebmaster

import (
	"fmt"
	"time"
)

// SearchQueryService - service for search query management
type SearchQueryService struct {
	client *Client
}

// newSearchQueryService - init SearchQueryService
func newSearchQueryService(cl *Client) *SearchQueryService {
	return &SearchQueryService{client: cl}
}

type SearchIndicator struct {
	TotalShows       float64 `json:"TOTAL_SHOWS"`
	TotalClicks      float64 `json:"TOTAL_CLICKS"`
	AvgShowPosition  float64 `json:"AVG_SHOW_POSITION"`
	AvgClickPosition float64 `json:"AVG_CLICK_POSITION"`
}

type PopularSearchQuery struct {
	QueryID    string          `json:"query_id"`
	QueryText  string          `json:"query_text"`
	Indicators SearchIndicator `json:"indicators"`
}

type PopularSeachQueryResponse struct {
	Queries  []*PopularSearchQuery `json:"queries"`
	DateFrom string                `json:"date_from"`
	DateTo   string                `json:"date_to"`
	Count    int                   `json:"count"`
}

type SeachAllHistoryIndicatorData struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type SearchAllHistoryIndicator struct {
	TotalShows  []*SeachAllHistoryIndicatorData `json:"TOTAL_SHOWS"`
	TotalClicks []*SeachAllHistoryIndicatorData `json:"TOTAL_Clicks"`
}

type SearchAllHistoryResponse struct {
	Indicators SearchAllHistoryIndicator `json:"indicators"`
}

type SearchSingleHistoryResponse struct {
	QueryID    string                    `json:"query_id"`
	QueryText  string                    `json:"query_text"`
	Indicators SearchAllHistoryIndicator `json:"indicators"`
}

// GetPopularSearchQueries - get popular queries, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-search-queries-popular.html
func (s *SearchQueryService) GetPopularSearchQueries(hostID string, dateFrom time.Time, dateTo time.Time, queryIndicator string, orderBy string, deviceTypeIndicator string, limit int, offset int) (PopularSeachQueryResponse, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	data["query_indicator"] = queryIndicator
	if orderBy != "" {
		data["order_by"] = orderBy
	} else {
		data["order_by"] = "TOTAL_SHOWS"
	}
	if deviceTypeIndicator != "" {
		data["device_type_indicator"] = deviceTypeIndicator
	} else {
		data["device_type_indicator"] = "ALL"
	}
	data["limit"] = limit
	data["offset"] = offset
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-queries/popular", s.client.userID, hostID)
	var result PopularSeachQueryResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// GetQueryAllHistory - get all query history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-search-queries-history-all.html
func (s *SearchQueryService) GetQueryAllHistory(hostID string, dateFrom time.Time, dateTo time.Time, queryIndicator string, deviceTypeIndicator string) (SearchAllHistoryResponse, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	data["query_indicator"] = queryIndicator
	if deviceTypeIndicator != "" {
		data["device_type_indicator"] = deviceTypeIndicator
	} else {
		data["device_type_indicator"] = "ALL"
	}
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-queries/all/history", s.client.userID, hostID)
	var result SearchAllHistoryResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// SearchSingleHistoryResponse - get sing search query history, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-search-queries-history.html
func (s *SearchQueryService) GetSingleSearchQueryHistory(hostID string, QueryID string, dateFrom time.Time, dateTo time.Time, queryIndicator string, deviceTypeIndicator string) (SearchSingleHistoryResponse, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	data["query_indicator"] = queryIndicator
	if deviceTypeIndicator != "" {
		data["device_type_indicator"] = deviceTypeIndicator
	} else {
		data["device_type_indicator"] = "ALL"
	}
	endpoint := fmt.Sprintf("user/%d/hosts/%s/search-queries/%s/history", s.client.userID, hostID, QueryID)
	var result SearchSingleHistoryResponse
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}
