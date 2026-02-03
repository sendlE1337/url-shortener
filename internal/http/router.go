package http_internal

import (
	shortener "SHORTNERED_URL/internal/service"
	"net/http"
)

// NewRouter создает HTTP роутер и регистрирует handler-ы
func NewRouter(s *shortener.Service) *http.ServeMux {
	mux := http.NewServeMux()

	//POST /shorten - создание короткой ссылки
	mux.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		HandlePost(w, r, s)
	})

	//Get /{id} - редирект на оригинальный URL
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleRedirect(w, r, s)
	})
	return mux
}
