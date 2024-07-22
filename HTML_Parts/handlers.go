package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSalutation(c *gin.Context) {

	c.String(http.StatusOK, "Salam Alikum")

}

func sendSalutationHtml(c *gin.Context) {

	c.HTML(http.StatusOK, "greeting.html", nil)
}

func sendSalutationHtml2(c *gin.Context) {

	name := c.Param("name") //give to name variable a value or url with tag name
	// example : when we make : home/salim ---> the value salim to the HTML page parametre

	c.HTML(http.StatusOK, "customgreeting.html", name)
}

func sendMany(c *gin.Context) {

	clubs := []string{"madira", "Sporting", "ManUtd", "Realmadrid", "Juventus", "Alnaasr"}

	c.HTML(http.StatusOK, "manydata.html", gin.H{
		"name":  "Ronaldo",
		"clubs": clubs,
	})
}

func getForm(c *gin.Context) {

	c.HTML(http.StatusOK, "form_u17.html", nil)
}

func postForm(c *gin.Context) {
	categ := "u17"
	name := c.PostForm("name")
	playpos := c.PostForm("playpos")

	c.HTML(http.StatusOK, "formResult.html", gin.H{
		"categ":    categ,
		"name":     name,
		"playepos": playpos,
	})
}

func getForm2(c *gin.Context) {

	c.HTML(http.StatusOK, "form_u19.html", nil)
}

func postForm2(c *gin.Context) {
	categ := "u19"
	name := c.PostForm("name")
	playpos := c.PostForm("playpos")

	c.HTML(http.StatusOK, "formResult.html", gin.H{
		"categ":    categ,
		"name":     name,
		"playepos": playpos,
	})

}
