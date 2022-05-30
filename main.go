package main

import (
	"github.com/fadhilradh/booker/handler"
	"github.com/fadhilradh/booker/router"
)

func main() {
	handler.StartDBConnection()
	router.Start()
}




