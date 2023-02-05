package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak-19/go-crud-postgres/database"
)

// create
func CreateBook(c *gin.Context) {
	var book *database.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}
