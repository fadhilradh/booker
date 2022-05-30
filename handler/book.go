package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fadhilradh/booker/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(ctx *gin.Context,) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Alhamdulillah",
	})
}

func GetBookByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func SearchHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	year := ctx.Query("year")
	publisher := ctx.Query("publisher")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"year": year,
		"publisher" : publisher,
	})
}



func PostNewBookHandler(ctx *gin.Context) {
	var newBook types.NewBook

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

func PostNewPublisherHandler(ctx *gin.Context) {
	var publisher types.Publisher

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