package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// Syntax : ---> toujours router.GET,or,POST (the url to handle it , function that handle)

	//send a simple message
	router.GET("/home", sendSalutation)
	// send an message inside html
	router.GET("/home/html", sendSalutationHtml)
	// send a customized html salutation  :name to pass the Hnadler the parametre that required
	router.GET("home/html/:name", sendSalutationHtml2)
	// send HTML page but with many data we send it as gin.H is MAP look at handlers.go
	router.GET("/ronaldo", sendMany)

	router.GET("/formu17", getForm)

	router.POST("/formu17", postForm)

	router.GET("/formu19", getForm2)

	router.POST("/formu19", postForm2)

	err := router.Run("localhost:8080")

	if err != nil {

		log.Fatal("bara amchi")
	}
}
