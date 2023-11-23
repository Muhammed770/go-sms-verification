package main

import (
	"github.com/Muhammed770/go-sms-verification/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	app.Routes()

	//run server
	router.Run(":8000")
}
