package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/books/:id", getBookByIdHandler)
	router.GET("/search", searchHandler)
	router.POST("/new-book", postNewBookHandler)
	router.POST("/new-publisher", postNewPublisherHandler)
	router.Run()
}

func rootHandler(ctx *gin.Context,) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Alhamdulillah",
	})
}

func getBookByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func searchHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	year := ctx.Query("year")
	publisher := ctx.Query("publisher")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"year": year,
		"publisher" : publisher,
	})
}

type NewBook struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func postNewBookHandler(ctx *gin.Context) {
	var newBook NewBook

	err:= ctx.ShouldBindJSON(&newBook)

	if err != nil {
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book added!",
		"title": newBook.Title,
		"author" : newBook.Author,
	})
}

type Publisher struct {
	Name string
	Location string
}

func postNewPublisherHandler(ctx *gin.Context) {
	var publisher Publisher

	err:= ctx.ShouldBindJSON(&publisher)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Publisher added!",
		"name": publisher.Name,
		"location" : publisher.Location,
	})
}