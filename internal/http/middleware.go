package http_internal

import "net/http"

//RequireMethod - проверяет http-метод запроса.
// В случае не совпадения return 405 - method not allowed
// в ином случа возов handler-a
func RequireMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
