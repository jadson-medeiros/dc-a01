package middleware

import "net/http"

func ContetTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PATCH,PUT,DELETE")
		next.ServeHTTP(w, r)
	})
}
