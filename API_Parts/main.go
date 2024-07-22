package main

import (
	"github.com/gin-gonic/gin"
)

type movie struct {
	ID    string `json : "id"`
	Title string `json : "title"`
	Price string `json : "price"`
}

var movies = []movie{
	{ID: "1", Title: "el hadj lakhder", Price: "500 da"},
	{ID: "2", Title: "Djam3i Family", Price: "800 da"},
	{ID: "3", Title: "Lawhama ", Price: "1000 da"},
}

func main() {

	router := gin.Default()

	router.GET("/listMov", showMovies)
	router.GET("/listMov/:id", showMoviesByID)
	router.POST("/listMov", Createmovie)
	router.PATCH("/listMov/changeprice/:id", Updatemovieprice)
	router.DELETE("listMov/delete/:id",DeleteByid)

	router.Run("localhost:8080")

}
