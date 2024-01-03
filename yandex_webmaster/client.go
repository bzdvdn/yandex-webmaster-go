package yandexwebmaster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

const (
	apiBaseUrl = "https://api.webmaster.yandex.net/v4/"
	YYYYMMDD   = "2006-01-02"
)

// YandexWebmaster Error
type YandexWebmasterError struct {
	HTTPCode     int
	Endpoint     string
	ErrorData    string
	ErrorMessage string
}

// Error returns string representation of the YandexWebmasterError
func (e *YandexWebmasterError) Error() string {
	return fmt.Sprintf("Http code: %d, endoint: %s, error data: %s, message: %s", e.HTTPCode, e.Endpoint, e.ErrorData, e.ErrorMessage)
}

// Client to interact with YandexWebmasterAPI
type Client struct {
	client       *http.Client
	token        string
	userID       int
	userIDLock   *sync.RWMutex
	Hosts        *HostService
	Sitemaps     *SitemapService
	Indexing     *IndexingService
	ImportantURL *ImportantURLService
	InsearchURL  *InsearchURLService
	Recrawl      *RecrawlService
	SearchQuery  *SearchQueryService
	Diagnostic   *DiagnosticService
}

// NewClient creates new Client to YandexWebmaster
func NewClient(token string) (*Client, error) {
	cl := &Client{
		client:     http.DefaultClient,
		token:      token,
		userID:     0,
		userIDLock: new(sync.RWMutex),
	}
	_, err := cl.getUserID()
	if err != nil {
		return nil, err
	}
	cl.Hosts = newHostService(cl)
	cl.Sitemaps = newSitemapService(cl)
	cl.Indexing = newIndexingService(cl)
	cl.ImportantURL = newImportURLService(cl)
	cl.InsearchURL = newInsearchURLService(cl)
	cl.Recrawl = newRecrawlService(cl)
	cl.SearchQuery = newSearchQueryService(cl)
	cl.Diagnostic = newDiagnosticService(cl)
	return cl, nil
}

// get user id for api requests
func (c *Client) getUserID() (int, error) {
	c.userIDLock.RLock()
	userID := c.userID
	c.userIDLock.RUnlock()
	if userID != 0 {
		return userID, nil
	}
	var responseData struct {
		UserID int `json:"user_id"`
	}
	endpoint := "user"
	_, err := c.sendAPIRequest(http.MethodGet, endpoint, nil, &responseData)
	if err != nil {
		return 0, err
	}

	c.userIDLock.Lock()
	c.userID = responseData.UserID
	userID = responseData.UserID
	c.userIDLock.Unlock()

	return userID, nil
}

// base method for api requests
func (c *Client) sendAPIRequest(method string, endpoint string, body interface{}, result interface{}) (*http.Response, error) {
	fullPath := apiBaseUrl + endpoint
	// fmt.Println(fullPath)
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}

	req, e := http.NewRequest(method, fullPath, buf)
	if e != nil {
		return nil, e
	}
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", c.token))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &YandexWebmasterError{http.StatusServiceUnavailable, endpoint, "", err.Error()}
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &YandexWebmasterError{resp.StatusCode, endpoint, string(respBody), err.Error()}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &YandexWebmasterError{resp.StatusCode, endpoint, string(respBody), ""}
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, &YandexWebmasterError{resp.StatusCode, endpoint, string(respBody), err.Error()}
	}

	return resp, nil
}

func (c *Client) generateURLWithGetParams(endpoint string, params map[string]interface{}) (string, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	query := url.Query()
	for param, value := range params {
		str := fmt.Sprintf("%v", value)
		query.Add(param, str)
	}
	url.RawQuery = query.Encode()
	return url.String(), nil
}

func (c *Client) makeGETRequestWithParams(endpoint string, params map[string]interface{}, result interface{}) (*http.Response, error) {
	endpoint, err := c.generateURLWithGetParams(endpoint, params)
	if err != nil {
		return nil, err
	}
	r, err := c.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return r, err
}
