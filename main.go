package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID  	 string `json:"id"`
    Title 	 string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust",Quantity: 7},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald",Quantity: 8},
	{ID: "3", Title: "The Art of War", Author: "Tsun Zou", Quantity: 12},
	{ID: "4", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 15},
}

func getBooks(c *gin.Context) {
c. IndentedJSON(http.StatusOK, books)
}

func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.Run("localhost:8080")
} 