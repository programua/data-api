package main

import (
	"data_api.com/data_api/handler"
	"fmt"
)

func main() {
	fmt.Println("main.go is run")

	router := handler.MakeRouter()
	router.Run()
}

