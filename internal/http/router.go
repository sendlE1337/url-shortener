package http_internal

import (
	shortener "SHORTNERED_URL/internal/service"
	"net/http"
)

// NewRouter создает HTTP роутер и регистрирует handler-ы
func NewRouter(s *shortener.Service) *http.ServeMux {
	mux := http.NewServeMux()

	//POST /shorten - создание короткой ссылки
	mux.HandleFunc("/shorten", RequireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		HandleShortener(w, r, s)
	}))

	//Get /{id} - редирект на оригинальный URL
	mux.HandleFunc("/", RequireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		HandleRedirect(w, r, s)
	}))
	return mux
}
