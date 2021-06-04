package main

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type game struct {
	player1 string
	player2 string
}

func socketEvents(io *socketio.Server) {
	users := make(map[string]string)
	games := make(map[string]game)

	io.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("new connection")
		return nil
	})

	io.OnEvent("/", "createGame", func(s socketio.Conn) string {
		id := createId()
		uid := s.ID()
		s.Join(id)
		users[uid] = id
		games[id] = game{player1: uid}

		return id
	})

	io.OnEvent("/", "joinGame", func(s socketio.Conn, id string) string {
		//checks if game exists
		if len(games[id].player1) == 0 {
			return "Error: Game not found"
		}

		//check if game is full
		if len(games[id].player2) != 0 {
			return "Error: Game is full"
		}

		//joins him to game
		users[s.ID()] = id
		games[id] = game{player1: games[id].player1, player2: s.ID()}
		io.BroadcastToRoom("/", id, "startGame")
		s.Join(id)

		return ""
	})
}
