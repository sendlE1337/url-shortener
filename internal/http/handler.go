package http_internal

import (
	"SHORTNERED_URL/internal/model"
	shortener "SHORTNERED_URL/internal/service"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// HandlerPost обрабатывает POST-запрос /shorten
func HandlePost(w http.ResponseWriter, r *http.Request, s *shortener.Service) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	originalURL := strings.TrimSpace(string(body))
	if originalURL == "" {
		http.Error(w, "empty url", http.StatusBadRequest)
		return
	}

	shortening, err := s.Shorten(string(body))
	if err != nil {
		if errors.Is(err, model.ErrAlreadyExists) {
			http.Error(w, "collision, try again", http.StatusConflict)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	response := struct {
		Identifier  string `json:"identifier"`
		OriginalURL string `json:"original_url"`
		ShortURL    string `json:"short_url"`
	}{
		Identifier:  shortening.Identifier,
		OriginalURL: shortening.OriginalURL,
		ShortURL:    "http://localhost:8080/" + shortening.Identifier,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// HandlerRedirect - обрабатывает GET /{id} и делает редирект
func HandleRedirect(w http.ResponseWriter, r *http.Request, s *shortener.Service) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	identifier := strings.TrimPrefix(r.URL.Path, "/")
	if identifier == "" {
		http.NotFound(w, r)
		return
	}

	originalURL, err := s.Redirect(identifier)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
