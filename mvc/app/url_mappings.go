package app

import "github.com/enginbulut/golang-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}