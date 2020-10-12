package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})

	if err := http.ListenAndServe(":1453", nil); err != nil {
		panic(err)
	}

}