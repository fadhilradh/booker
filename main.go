package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Price int `json:"price" binding:"required,number"`
}

func postNewBookHandler(ctx *gin.Context) {
	var newBook NewBook

	err:= ctx.ShouldBindJSON(&newBook)

	if err != nil {
		errorMessages := []string{}
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else {
			fmt.Println(err)
			errorMessages = append(errorMessages, err.Error())
		}
	
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book added!",
		"title": newBook.Title,
		"author" : newBook.Author,
		"price" : newBook.Price,
	})
	return
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