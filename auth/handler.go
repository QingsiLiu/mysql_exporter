package auth

import (
	"mysql_exporter/config"
	"net/http"
)

func BasicAuth(config *config.AuthConfig, handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		secret := request.Header.Get("authorization")
		if isAuth(secret, config) {
			handler.ServeHTTP(writer, request)
		} else {
			writer.Header().Set("www-authenticate", `basic realm="my site"`)
			writer.WriteHeader(http.StatusUnauthorized)
		}
	}
}
