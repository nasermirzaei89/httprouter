# HTTP Router
Simple Golang HTTP Router

## Sample
```go
package main

import (
	"github.com/nasermirzaei89/httprouter"
	"log"
	"net/http"
)

func main() {
	h := httprouter.New()

	h.Get("^/ping$", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe("0.0.0.0:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
```
