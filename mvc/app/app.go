package app

import (
	"github.com/enginbulut/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":1299", nil); err != nil {
		panic(err)
	}
}