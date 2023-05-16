package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Service(address, port string) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	Route(g)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", address, port),
		Handler: g,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}
}
