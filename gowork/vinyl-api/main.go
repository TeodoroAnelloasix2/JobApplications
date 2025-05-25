package main

import (
	"fmt"
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artista string  `json:"artist"`
	Price   float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Nevermind", Artista: "Nirvana", Price: 12.99},
	{ID: "2", Title: "Back in Black", Artista: "AC/DC", Price: 10.50},
	{ID: "3", Title: "The Dark Side of the Moon", Artista: "Pink Floyd", Price: 15.00},
	{ID: "4", Title: "Thriller", Artista: "Michael Jackson", Price: 14.99},
	{ID: "5", Title: "Abbey Road", Artista: "The Beatles", Price: 13.50},
	{ID: "6", Title: "Hotel California", Artista: "Eagles", Price: 11.75},
}

func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error, no se pudo agregar la entrada"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})

}

func deleteAlbumbyId(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			albums = slices.Delete(albums, i, i+1)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})

}

func modifyById(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos, no se pudo procesar la solicitud"})
		return
	}
	id := c.Param("id")
	for i, album := range albums {
		if album.ID == id {
			albums[i].Title = newAlbum.Title
			albums[i].Artista = newAlbum.Artista
			albums[i].Price = newAlbum.Price
			c.IndentedJSON(http.StatusCreated, albums[i])
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {

	fmt.Println("Bienvendio a Vinyl-API")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola Mundo",
		})
	})

	router.GET("/albums", getAlbums)
	router.POST("/addAlbum", addAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumbyId)
	router.PUT("/modify/:id", modifyById)
	router.Run("localhost:8080")
}
