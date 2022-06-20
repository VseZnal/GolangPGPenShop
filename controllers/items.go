package controllers

import (
	"GolangPGPenShop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /items
// Get all items

func FindItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})
}

//POST /items
//Create new item

func CreateItemNew(c *gin.Context) {
	// Validate input
	var input models.CreateItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create item
	item := models.Item{Title: input.Title, Description: input.Description, Contact: input.Contact}
	models.DB.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// GET /items/:id
// Find a item

func FindItem(c *gin.Context) {
	var item models.Item

	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}
