package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/siddoinghisjob/urlShortener/utils"
)

func setupAPI() {
	defer utils.Client.Close()
	if utils.DBErr != nil {
		panic(utils.DBErr)
	}

	defer utils.DB.Close()
	gin.SetMode(gin.ReleaseMode)

	ticker := time.NewTicker(10 * time.Second)
	utils.Logging(ticker)
	defer ticker.Stop()

	router := gin.Default()
	router.ForwardedByClientIP = true

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/healthz", utils.Healthz)
	router.POST("/url", utils.Post)
	router.GET("/url/:id", utils.Get)
	router.GET("/analytics/:id", utils.Analytics)

	router.Run()
}
