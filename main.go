package main

import (
	"github.com/gin-gonic/gin"
	models "lowcode/models"
	"net/http"
)

type CreateSnippetInput struct {
	Title        string `json:"title" binding:"required"`
	Code         string `json:"code" binding:"required"`
	IsActive     bool   `json:"isActive"`
	PlatformType int    `json:"platformType" binding:"required"`
}

func getAllSnippets(context *gin.Context) {
	var snippets []models.Snippet
	models.ConnectDB().Find(&snippets)

	context.JSON(http.StatusOK, gin.H{
		"snippets": snippets,
	})
}

func createSnippet(context *gin.Context) {
	var input CreateSnippetInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	snippet := models.Snippet{
		Title:        input.Title,
		Code:         input.Code,
		IsActive:     input.IsActive,
		PlatformType: input.PlatformType,
	}
	models.ConnectDB().Create(&snippet)

	context.JSON(http.StatusOK, gin.H{
		"snippet": snippet,
	})
}

func getSnippetById(context *gin.Context) {
	var snippet models.Snippet
	models.ConnectDB().First(&snippet, context.Param("id"))
	if snippet.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Row doesnt exist"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"snippet": snippet})
}

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/api/snippets", getAllSnippets)
	route.GET("/api/snippets/:id", getSnippetById)
	route.POST("/api/snippets", createSnippet)

	route.Run(":5000")
}
