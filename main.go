package main

import (
	"github.com/CuriousHet/Notify/notification"
	"github.com/CuriousHet/Notify/server"
)

func main() {
	queue := notification.NewQueue(100)
	dispatcher := notification.NewDispatcher(queue, 3)
	dispatcher.Start(5) // Start 5 workers

	server.StartGRPCServer(queue)
}
