package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()

	if err := http.ListenAndServe(":1299", nil); err != nil {
		panic(err)
	}
}