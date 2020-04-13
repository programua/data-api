package main

import (
	"github.com/programua/data-api/handler"
	"net/http"
)

func main() {
	http.Handle("/", MakeRouter())
}

