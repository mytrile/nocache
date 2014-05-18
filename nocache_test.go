package nocache

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-martini/martini"
)

func Test_UpdateCacheHeadersInProductionEnv(t *testing.T) {
	os.Setenv("MARTINI_ENV", "production")
	recorder := httptest.NewRecorder()
	m := martini.New()

	m.Use(UpdateCacheHeaders())
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)
	if len(recorder.HeaderMap) != 0 {
		t.Error("Unexpected headers were added")
	}
}

func Test_UpdateCacheHeadersInDevelopmentEnv(t *testing.T) {
	responseHeaders := map[string]string{
		"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
		"Pragma":        "no-cache",
		"Expires":       "Fri, 29 Aug 1997 02:14:00 EST",
	}

	os.Setenv("MARTINI_ENV", "development")
	recorder := httptest.NewRecorder()
	m := martini.New()

	m.Use(UpdateCacheHeaders())
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)

	for key, value := range responseHeaders {
		if recorder.HeaderMap[key][0] != value {
			t.Error("Missing header: %s", key)
		}
	}
}
