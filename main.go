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

func main() {
	sockets, _ := socketio.NewServer(nil)
	sockets.On("connection", connection)
	log.Println("Listening at http://localhost:8000")
	http.ListenAndServe(":8000", http.FileServer(http.Dir("ui/dist")))
	death.NewDeath(SYS.SIGINT, SYS.SIGTERM).WaitForDeath()
}
