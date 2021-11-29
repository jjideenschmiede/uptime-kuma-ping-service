//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of uptime-kuma-ping-service.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

// Response is to return a response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {

	// Set gin mode
	gin.SetMode(gin.ReleaseMode)

	// Disable logging
	gin.DefaultWriter = io.Discard

	// Define router
	router := gin.New()

	// Add ping pong
	router.GET("/ping", func(c *gin.Context) {

		// Log to console
		log.Println("We played a little ping pong.")

		// Return data
		c.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "pong",
		})

	})

	// Check ssl & start server
	if os.Getenv("ENABLE_SSL") == "true" {
		router.RunTLS(":443", "/go/src/app/files/certificates/"+os.Getenv("CERTIFICATE_CRT_NAME"), "/go/src/app/files/certificates/"+os.Getenv("CERTIFICATE_KEY_NAME"))
	} else {
		router.Run(":443")
	}

}
