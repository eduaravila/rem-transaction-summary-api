package server

import (
	"log"
	"net/http"
	"strings"
)

func RunHTTPServer(prefix, addrs string, createHandler func() http.Handler) {

	router := http.NewServeMux()

	router.Handle(prefix+"/", http.StripPrefix(prefix, createHandler()))

	withCors(*router)

	log.Fatal(http.ListenAndServe(addrs, router))
}

func withCors(mux http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessControlAllowHeaders strings.Builder
		accessControlAllowHeaders.WriteString("Accept,")
		accessControlAllowHeaders.WriteString("Content-Type,")
		accessControlAllowHeaders.WriteString("Content-Length,")
		accessControlAllowHeaders.WriteString("Accept-Encoding,")
		accessControlAllowHeaders.WriteString("X-CSRF-Token,")
		accessControlAllowHeaders.WriteString("Authorization")

		r.Header = map[string][]string{
			"Access-Control-Allow-Origin":      {"http://localhost"},
			"Access-Control-Allow-Credentials": {"true"},
			"Access-Control-Allow-Headers":     {accessControlAllowHeaders.String()},
		}

		mux.ServeHTTP(w, r)
	})
}
