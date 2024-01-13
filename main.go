package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func main() {
	e := echo.New()

	// Add middleware to inject the manager into the context
	mgr := NewWebsocketManager()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("ws", mgr)
			return next(c)
		}
	})

	e.POST("/trigger", trigger)
	e.GET("/ws", ws)
	e.Static("/*", "public")
	e.Logger.Fatal(e.Start(":9090"))
}

func ws(c echo.Context) error {
	mgr := c.Get("ws").(*WebsocketManager)

	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		mgr.AddConnection(conn)
		defer mgr.RemoveConnection(conn)

		fmt.Println("client connected")
		hello := Message{"hello", "world"}
		websocket.JSON.Send(conn, hello)
		for {
			msg := Message{}
			err := websocket.JSON.Receive(conn, &msg)
			if err != nil {
				continue
			}
			fmt.Printf("Received: %#v\n", msg)
			// websocket.JSON.Send(conn, msg)
			mgr.SendMessage(msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

// Broabcast a message to all connected clients
func trigger(c echo.Context) error {
	mgr := c.Get("ws").(*WebsocketManager)
	mgr.SendMessage(Message{"trigger", "triggered"})

	return c.HTML(200, "<li>Triggered!</li>")
}
