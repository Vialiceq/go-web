package routers

import (
	//"1/db"
	h "1/controller/handles"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func Routers(url *string) {

	router := gin.Default()
	router.GET("/albums", h.GetAlbums)
	router.GET("/albums/:id", h.GetAlbumByID)
	router.POST("/albums", h.PostAlbums)
	router.Run(*url)
}
