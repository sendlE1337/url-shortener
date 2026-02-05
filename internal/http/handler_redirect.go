package http_internal

import (
	"SHORTNERED_URL/internal/model"
	shortener "SHORTNERED_URL/internal/service"
	"errors"
	"net/http"
	"strings"
)

// HandlerRedirect - обрабатывает GET /{id} и делает редирект
func HandleRedirect(w http.ResponseWriter, r *http.Request, s *shortener.Service) {
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
