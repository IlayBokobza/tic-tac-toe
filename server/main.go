package main

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	io := socketio.NewServer(nil)

	//create events
	socketEvents(io)

	//sets up socket.io
	go io.Serve()
	defer io.Close()
	http.Handle("/socket.io/", io)

	//app route
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
