package swen

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("query")
	limitQ := r.URL.Query().Get("limit")
	if limitQ == "" {
		limitQ = "5"
	}
	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		http.Error(w, "invalid limit", http.StatusBadRequest)
		return
	}
	news, err := c.service.GetNewsForQuery(r.Context(), query, limit)
	if err != nil {
		log.Printf("Error getting news: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
