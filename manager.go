package main

import (
	"sync"

	"golang.org/x/net/websocket"
)

type WebsocketManager struct {
	connections map[*websocket.Conn]bool
	lock        sync.Mutex
}

func NewWebsocketManager() *WebsocketManager {
	return &WebsocketManager{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (manager *WebsocketManager) AddConnection(conn *websocket.Conn) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.connections[conn] = true
}

func (manager *WebsocketManager) RemoveConnection(conn *websocket.Conn) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	delete(manager.connections, conn)
}

func (manager *WebsocketManager) SendMessage(message interface{}) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	for conn := range manager.connections {
		err := websocket.JSON.Send(conn, message)
		if err != nil {
			// Handle error, possibly remove the connection
			delete(manager.connections, conn)
		}
	}
}
