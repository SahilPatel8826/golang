package main

import (
	api "chatapp2/api"
	store "chatapp2/connectdb"
	ws "chatapp2/websocket"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	store.ConnectDB()

	// Auto-migrate tables
	store.DB.AutoMigrate(&ws.OutboundMessage{})
	hub := ws.NewHub()

	go hub.Run()
	e := echo.New()
	api.RegisterRoute(e, hub)

	fmt.Println("Server starting on :9000 ...")
	// start your websocket hub / routes here
	e.Logger.Fatal(e.Start(":9000"))
}
