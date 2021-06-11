package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	socketio "github.com/googollee/go-socket.io"
	"github.com/ilaybokobza/tic-tac-toe/server/games"
	"github.com/ilaybokobza/tic-tac-toe/server/users"
)

func main() {
	io := socketio.NewServer(nil)

	//resets data files
	err := users.ResetFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = games.ResetFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//create events
	socketEvents(io)

	//sets up socket.io
	go io.Serve()
	// defer io.Close()
	http.Handle("/socket.io/", io)

	//app route
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
