package main

import (
	"008_original_api"
	"net/http"
)

func main() {
	http.Handle("/", handler.MakeRouter())
}

