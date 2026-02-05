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
func HandleShortener(w http.ResponseWriter, r *http.Request, s *shortener.Service) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	originalURL := strings.TrimSpace(string(body))
	if err := validateURL(originalURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
