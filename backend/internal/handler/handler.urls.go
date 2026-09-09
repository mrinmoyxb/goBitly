package handler

import (
	"context"
	"encoding/json"
	"goBitly/internal/model"
	"goBitly/internal/services"
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
)

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

func (handler *URLHandler) CheckHealth(w http.ResponseWriter, r *http.Request){
	msg := "hello from goBitly!!!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func (handler *URLHandler) CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateShortURLRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserId == 0 {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}
	if req.OriginalURL == "" {
		http.Error(w, "original url is required", http.StatusBadRequest)
		return
	}

	url, err := handler.service.CreateShortURLService(r.Context(), req.UserId, req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

func (handler *URLHandler) GetByShortURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	if shortURL == "" {
		http.Error(w, "short url is required", http.StatusBadRequest)
		return
	}

	url, err := handler.service.GetURLByShortURLService(r.Context(), shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}

func (handler *URLHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	if shortURL == "" {
		http.Error(w, "short url is required", http.StatusBadRequest)
		return
	}

	url, err := handler.service.GetURLByShortURLService(r.Context(), shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)

	go func(){
		ctx := context.Background()
		if _, err := handler.service.IncrementClickCountService(ctx, url.ShortURL); err != nil {
			log.Printf("failed to increment click count for %s: %v", shortURL, err)
		}
	}()
}

func (handler *URLHandler) GetByOriginalURLHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("original_url")

	if originalURL == "" {
		http.Error(w, "original url is required", http.StatusBadRequest)
		return
	}

	url, err := handler.service.GetURLByOriginalURLService(r.Context(), originalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}

func (handler *URLHandler) DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	if shortURL == "" {
		http.Error(w, "short url is required", http.StatusBadRequest)
		return
	}

	success, err := handler.service.DeleteURLService(r.Context(), shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Success bool `json:"success"`
	}{
		Success: success,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (handler *URLHandler) GetClickCountHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	if shortURL == "" {
		http.Error(w, "short url is required", http.StatusBadRequest)
		return
	}

	count, err := handler.service.GetClickCountsService(r.Context(), shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response := struct {
		Count int64 `json:"count"`
	}{
		Count: count,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
