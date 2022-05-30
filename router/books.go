package router

import (
	"github.com/fadhilradh/booker/handler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	v1 := router.Group("/v1")
	
	v1.GET("/", handler.RootHandler)
	v1.GET("/books", handler.GetBooks)
	v1.GET("/books/:id", handler.GetBookByIdHandler)
	v1.GET("/search", handler.SearchHandler)
	v1.POST("/new-book", handler.PostNewBookHandler)
	v1.POST("/new-publisher", handler.PostNewPublisherHandler)
	router.Run()	
}

