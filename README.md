# HTTP Router
Simple Golang HTTP Router

[![Build Status](https://travis-ci.org/nasermirzaei89/httprouter.svg?branch=master)](https://travis-ci.org/nasermirzaei89/httprouter)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/httprouter)](https://goreportcard.com/report/github.com/nasermirzaei89/httprouter)
[![GoDoc](https://godoc.org/github.com/nasermirzaei89/httprouter?status.svg)](https://godoc.org/github.com/nasermirzaei89/httprouter)
[![GitHub license](https://img.shields.io/github/license/nasermirzaei89/httprouter.svg)](https://github.com/nasermirzaei89/httprouter/blob/master/LICENSE)


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
