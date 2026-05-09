package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Print(start)

		next.ServeHTTP(w, req)
		
		end := time.Now()
		log.Print(end)
	})
}
