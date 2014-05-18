package nocache

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

var responseHeaders = map[string]string{
	"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
	"Pragma":        "no-cache",
	"Expires":       "Fri, 29 Aug 1997 02:14:00 EST",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// UpdateCacheHeaders removes ETag related headers and injects regular "no cache" headers
func UpdateCacheHeaders() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {

		if os.Getenv("MARTINI_ENV") == "development" {
			for _, value := range etagHeaders {
				if r.Header.Get(value) != "" {
					fmt.Println(r.Header.Get(value))
					r.Header.Del(value)
				}
			}
			headers := w.Header()
			for key, value := range responseHeaders {
				headers.Set(key, value)
			}
		}
		return
	}
}
