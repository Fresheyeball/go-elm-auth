package main

import (
	"log"
	"net/http"
	"strconv"
	SYS "syscall"

	"github.com/googollee/go-socket.io"
	"github.com/vrecan/death"
)

const (
	messageKey = "message"
	roomKey    = "room"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func connection(socket socketio.Socket) {
	log.Println("Connected!")
	socket.Join(roomKey)
}

func message(sockets *socketio.Server) func(string) {
	return func(message string) {
		sockets.BroadcastTo(
			roomKey,
			messageKey,
			strconv.Itoa(factorial(len(message))))
	}
}

func newSocket() *socketio.Server {
	sockets := attemptGet(socketio.NewServer(nil)).(*socketio.Server)
	attempt(sockets.On("connection", connection))
	attempt(sockets.On(messageKey, message(sockets)))
	return sockets
}

func route() {
	http.Handle("/socket.io/", newSocket())
	http.Handle("/", http.FileServer(http.Dir("ui/dist")))
}

func main() {
	route()
	log.Println("Listening at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath()
}
