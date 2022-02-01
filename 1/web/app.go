package web

import (
	/*
		"1/db"
		"1/model"*/
	"github.com/gin-gonic/gin"
	"net/http"

	//c "1/handles"
	r "1/controller/routers"
)

// albums slice to seed record album data.
/*
type App0 struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}
*/
func App() {

	url := "localhost:9095"

	r.Routers(&url)
}
