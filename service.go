package swen

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Service struct {
	apiKey         string
	searchEngineID string
}

func NewService(apiKey, searchEngineID string) *Service {
	return &Service{
		apiKey:         apiKey,
		searchEngineID: searchEngineID,
	}
}

func (s *Service) GetNewsForQuery(ctx context.Context, query string, limit int) ([]NewsItem, error) {
	query = strings.TrimSpace(query) + " news"
	searchURL := "https://www.googleapis.com/customsearch/v1?key=" + s.apiKey + "&cx=" + s.searchEngineID + "&q=" + url.QueryEscape(query) + "&num=10&siteSearch=moneycontrol.com"
	resp, err := http.Get(searchURL)
	if err != nil {
		log.Printf("Request error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response read error: %v", err)
		return nil, err
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		log.Printf("Response parse error: %v", err)
		return nil, err
	}
	return getNewsItemsFromResponse(responseMap), nil
}

func getNewsItemsFromResponse(responseMap map[string]interface{}) []NewsItem {
	if responseMap["items"] == nil {
		return nil
	}
	items := responseMap["items"].([]interface{})
	newsItems := make([]NewsItem, 0, len(items))
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		if itemMap["title"] == nil || itemMap["title"].(string) == "" || itemMap["link"] == nil || itemMap["link"].(string) == "" || itemMap["pagemap"] == nil || itemMap["displayLink"] == nil || itemMap["formattedUrl"] == nil {
			continue
		}
		newsItems = append(newsItems, NewsItem{
			Title:     itemMap["title"].(string),
			URL:       itemMap["link"].(string),
			Image:     itemMap["pagemap"].(map[string]interface{})["cse_image"].([]interface{})[0].(map[string]interface{})["src"].(string),
			Source:    itemMap["displayLink"].(string),
			SourceURL: itemMap["formattedUrl"].(string),
		})
	}
	return newsItems
}
