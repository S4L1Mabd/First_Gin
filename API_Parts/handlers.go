package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showMovies(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, movies)

}
func showMoviesByID(c *gin.Context) {
	id := c.Param("id")
	var index int

	for k, v := range movies {

		if id == v.ID {
			index = k
		}
	}

	c.IndentedJSON(http.StatusOK, movies[index])
}

func Createmovie(c *gin.Context) {
	var newmovie movie

	err := c.BindJSON(&newmovie)

	if err != nil {

		fmt.Println("error bara ")
	}

	movies = append(movies, newmovie)

	c.IndentedJSON(http.StatusOK, movies)
}

func Updatemovieprice(c *gin.Context) {

	var index int

	id := c.Param("id")

	for ind, v := range movies {

		if id == v.ID {

			index = ind
		}
	}
	movies[index].Price = "600 da"
	c.IndentedJSON(http.StatusOK, movies[index])
}

func DeleteByid(c *gin.Context) {

	var index int

	id := c.Param("id")

	for ind, v := range movies {

		if id == v.ID {

			index = ind
		}
	}

	movies = append(movies[:index], movies[index+1:]...)
	c.IndentedJSON(http.StatusOK, nil)

}
