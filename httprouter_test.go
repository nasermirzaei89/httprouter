package httprouter_test

import (
	"fmt"
	. "github.com/nasermirzaei89/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ServeHTTP(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	h := New()

	h.Get("^/ping$", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status code: '%d' got '%d'", http.StatusOK, status)
	}
}

func BenchmarkHandler_ServeHTTP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			b.Fatal(err)
		}

		rec := httptest.NewRecorder()
		h := New()

		for j := 0; j < 100; j++ {
			h.Get(fmt.Sprintf("^/route%d$", j), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
		}

		h.Get("^/ping$", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		h.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusOK {
			b.Errorf("expected status code: '%d' got '%d'", http.StatusOK, status)
		}
	}
}
