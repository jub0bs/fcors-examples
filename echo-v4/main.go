package main

import (
	"log"
	"net/http"

	"github.com/jub0bs/fcors"
	"github.com/labstack/echo/v4"
)

func main() {
	// create an Echo instance
	e := echo.New()

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

	// apply the CORS middleware to your Echo instance
	e.Use(echo.WrapMiddleware(cors))

	// register the hello handler for the /hello route
	e.GET("/hello", hello)

	// start the server on port 8080
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal()
	}
}

func hello(c echo.Context) error {
	data := struct {
		Hello string `json:"hello"`
	}{
		Hello: "world",
	}
	return c.JSON(http.StatusOK, data)
}
