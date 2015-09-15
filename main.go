package main

import (
	"log"
	"net/http"
	SYS "syscall"

	"github.com/googollee/go-socket.io"
	"github.com/vrecan/death"
)

func connection(socket socketio.Socket) {
	log.Println("Connected!")
}

func newSocket() *socketio.Server {
	sockets := attemptGet(socketio.NewServer(nil)).(*socketio.Server)
	attempt(sockets.On("connection", connection))
	return sockets
}

func route() {
	http.Handle("/socket/", newSocket())
	http.Handle("/", http.FileServer(http.Dir("ui/dist")))
}

func main() {
	route()
	log.Println("Listening at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath()
}
