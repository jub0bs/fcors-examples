package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jub0bs/fcors"
)

func main() {
	// create an engine
	engine := gin.Default()

	// configure the CORS middleware
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("https://example.com"),
		fcors.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		),
		fcors.WithRequestHeaders("Authorization"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware to the engine
	engine.Use(adapter.Wrap(cors))

	// register the hello handler for the /hello route
	engine.GET("/hello", hello)

	// start the server on port 8080
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
}
