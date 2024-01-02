package yandexwebmaster

import (
	"fmt"
	"net/http"
	"time"
)

// RecrawlService - service for manage recrawl tasks
type RecrawlService struct {
	client *Client
}

// newRecrawlService func for init RecrawlService
func newRecrawlService(cl *Client) *RecrawlService {
	return &RecrawlService{client: cl}
}

// RecrawlURLResponse ...
type RecrawlURLResponse struct {
	TaskID string `json:"task_id"`
	URL    string `json:"url"`
}

type RecrawlTask struct {
	TaskID    string `json:"task_id"`
	URL       string `json:"url"`
	AddedTime string `json:"added_time"`
	State     string `json:"state"`
}

type RecrawlTasks struct {
	Tasks []*RecrawlTask `json:"tasks"`
}

type RecrawlQuota struct {
	DailyQuota     int `json:"daily_quota"`
	QuotaRemainder int `json:"quota_remainder"`
}

// start recrawl url, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-recrawl-post.html
func (s *RecrawlService) RecrawlURL(hostID string, url string) (RecrawlURLResponse, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/recrawl/queue", s.client.userID, hostID)
	data := make(map[string]interface{})
	data["url"] = url
	var result RecrawlURLResponse
	_, err := s.client.sendAPIRequest(http.MethodPost, endpoint, data, &result)
	return result, err
}

// get recrawl task, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-recrawl-task-get.html
func (s *RecrawlService) GetRecrawlTask(hostID string, taskID string) (RecrawlTask, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/recrawl/queue/%s", s.client.userID, hostID, taskID)
	var result RecrawlTask
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

// get recrawl tasks, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-recrawl-get.html
func (s *RecrawlService) GetRecrawlTasks(hostID string, dateFrom time.Time, dateTo time.Time, limit int, offset int) (RecrawlTasks, error) {
	data := make(map[string]interface{})
	data["date_from"] = dateFrom.Format(YYYYMMDD)
	data["date_to"] = dateTo.Format(YYYYMMDD)
	data["limit"] = limit
	data["offset"] = limit
	endpoint := fmt.Sprintf("user/%d/hosts/%s/recrawl/queue", s.client.userID, hostID)
	var result RecrawlTasks
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get recrawl quota, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-recrawl-quota-get.html
func (s *RecrawlService) GetRecrawlQuota(hostID string) (RecrawlQuota, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/recrawl/quota", s.client.userID, hostID)
	var result RecrawlQuota
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}
