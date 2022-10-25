package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	h := NewHandler()
	e := echo.New()
	e.GET("/websocket", h.WebSocketHandler)
	log.Fatal(e.Start("0.0.0.0:8080"))
}
