package yandexwebmaster

import (
	"fmt"
	"net/http"
)

// Sitemap service for sitemap management
type SitemapService struct {
	client *Client
}

// func for init sitemap service
func newSitemapService(cl *Client) *SitemapService {
	return &SitemapService{client: cl}
}

type Sitemap struct {
	SitemapID      string `json:"sitemap_id"`
	SitemapURL     string `json:"sitemap_url"`
	LastAccessDate string `json:"last_access_date"`
}

type Sitemaps struct {
	Sitemaps []*Sitemap `json:"sitemaps"`
}

type AddedUserSitemap struct {
	SitemapID  string `json:"sitemap_id"`
	SitemapURL string `json:"sitemap_url"`
	AddedDate  string `json:"added_date"`
}

type AddedSitemap struct {
	SitemapID string `json:"sitemap_id"`
}

// get sitemaps, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-sitemaps-get.html
func (s *SitemapService) GetSitemaps(hostID string, limit int, parentID string, fromSiteID string) (Sitemaps, error) {
	data := make(map[string]interface{})
	if limit != 0 {
		data["limit"] = limit
	} else {
		data["limit"] = 10
	}
	if parentID != "" {
		data["parent_id"] = parentID
	}
	if fromSiteID != "" {
		data["from_site_id"] = fromSiteID
	}
	var result Sitemaps
	endpoint := fmt.Sprintf("user/%d/hosts/%s/sitemaps", s.client.userID, hostID)
	_, err := s.client.makeGETRequestWithParams(endpoint, data, &result)
	return result, err
}

// get site map, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-sitemaps-sitemap-id-get.html
func (s *SitemapService) GetSitemap(hostID string, sitemapID string) (Sitemap, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/sitemaps/%s", s.client.userID, hostID, sitemapID)
	var result Sitemap
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

// get user added sitemaps, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-user-added-sitemaps-sitemap-id-get.html
func (s *SitemapService) GetUserAddedSitemap(hostID string, sitemapID string) (AddedUserSitemap, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/user-added-sitemaps/%s", s.client.userID, hostID, sitemapID)
	var result AddedUserSitemap
	_, err := s.client.sendAPIRequest(http.MethodGet, endpoint, nil, &result)
	return result, err
}

// add sitemap, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-user-added-sitemaps-post.html
func (s *SitemapService) AddSitemap(hostID string, url string) (AddedSitemap, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/user-added-sitemaps", s.client.userID, hostID)
	var result AddedSitemap
	data := make(map[string]interface{})
	data["url"] = url
	_, err := s.client.sendAPIRequest(http.MethodPost, endpoint, data, &result)
	return result, err
}

// Delete sitemap, doc: https://yandex.ru/dev/webmaster/doc/dg/reference/host-user-added-sitemaps-sitemap-id-delete.html
func (s *SitemapService) DeleteSitemap(hostID string, sitemapID string) (interface{}, error) {
	endpoint := fmt.Sprintf("user/%d/hosts/%s/user-added-sitemaps/%s", s.client.userID, hostID, sitemapID)
	var result interface{}
	_, err := s.client.sendAPIRequest(http.MethodDelete, endpoint, nil, &result)
	return result, err
}
