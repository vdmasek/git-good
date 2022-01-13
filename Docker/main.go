package main

import (
	"log"
	"net/http"

	"github.com/docker/util"

	"github.com/gin-gonic/gin"
)

var config *util.Config

func init() {
	configur, err := util.LoadConfig(".")

	if err != nil {
		print("EROARE", err)
		log.Fatal("cannot load config err", err)
	}
	config = &configur
	util.LoadCollection("./static/")
}

func main() {
	var router *gin.Engine

	router = gin.Default()

	router.Static("/static/", "./static")

	router.LoadHTMLGlob("templates/*")
	router.GET("/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"products": util.PCollection.Listaa,
		})
	})

	//router.Run()

	router.Run(config.ServerAddress + ":" + config.AppPort)
}
