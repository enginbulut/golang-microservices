package app

import (
	"github.com/enginbulut/golang-microservices/src/api/controllers/polo"
	"github.com/enginbulut/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
