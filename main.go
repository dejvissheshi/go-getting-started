package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello, this is the test endpoint!")
	})

	router.GET("/calculate/:id", CalculatePackages)

	router.GET("/rollback", RollbackPackageChanges)
	router.GET("/add/:id", AddNewPackages)
	router.GET("/remove/:id", RemovePackages)

	router.Run(":" + port)
}
