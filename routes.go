package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/file"
	"net/http"
	"strconv"
)

var defaultItems = []int{250, 500, 1000, 2000, 5000}

// RollbackPackageChanges is a handler for the rollback endpoint
func RollbackPackageChanges(c *gin.Context) {
	err := file.RollbackFileToInitialState("data.csv", defaultItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numbers": numbers})
}

// AddNewPackages is a handler for the add endpoint
func AddNewPackages(c *gin.Context) {
	// Extract the "id" parameter from the URL

	id := c.Param("id")

	newPackageSize, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = file.AddNumbersToCSV("data.csv", []int{newPackageSize})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numbers": numbers})
}

// RemovePackages is a handler for the delete endpoint
func RemovePackages(c *gin.Context) {
	// Extract the "id" parameter from the URL
	id := c.Param("id")

	newPackageSize, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = file.DeleteNumbersFromCSV("data.csv", []int{newPackageSize})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numbers": numbers})
}

// CalculatePackages is a handler for the Calculate endpoint
func CalculatePackages(c *gin.Context) {
	// Extract the "id" parameter from the UR
	id := c.Param("id")

	itemsOrdered, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	packages := Calculate(itemsOrdered, defaultItems)
	// Create a response JSON
	response := packages

	c.JSON(http.StatusOK, response)
}
