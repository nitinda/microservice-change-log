package middlewares

import (
	"net/http"

	"github.com/nitinda/microservice-change-log/logger"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(rw, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next(rw, r)
	}
}
