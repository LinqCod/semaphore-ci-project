package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	path, err := filepath.Abs("")
	if err != nil {
		log.Fatal(err)
	}
	router.LoadHTMLGlob(path + "/templates/*")

	router.GET("/", showIndexPage)

	router.Run(":4000")
}
