package urlShortener

import (
	"net/http"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(writer, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(writer, request)
	}
}
