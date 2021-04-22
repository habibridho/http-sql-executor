package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	server := echo.New()
	server.GET("ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "pong",
		})
	})

	server.Logger.Fatal(server.Start(":3000"))
}
