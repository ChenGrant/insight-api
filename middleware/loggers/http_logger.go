package loggers

import (
	"log"
	"net/http"
)

type HTTPLogger struct {
	http.ResponseWriter
	statusCode int
}

func HTTPLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lw := &HTTPLogger{ResponseWriter: w, statusCode: http.StatusOK}
		next(lw, r)
		log.Println(r.Method, r.URL.Path, lw.statusCode)
	}
}
