package main

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type Game struct {
	player1 string
	player2 string
	turn    int
}

func socketEvents(io *socketio.Server) {
	users := make(map[string]string)
	games := make(map[string]Game)

	io.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("new connection")
		return nil
	})

	io.OnEvent("/", "createGame", func(s socketio.Conn) string {
		id := createId()
		uid := s.ID()
		s.Join(id)
		users[uid] = id
		games[id] = Game{player1: uid}

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
		games[id] = Game{player1: games[id].player1, player2: s.ID(), turn: 1}
		io.BroadcastToRoom("/", id, "startGame")
		s.Join(id)

		return ""
	})

	io.OnEvent("/", "madeTurn", func(s socketio.Conn, cords []int) string {
		uid := s.ID()
		gameId := users[uid]
		game := games[gameId]
		var playerType int

		if uid == game.player1 {
			playerType = 1
		} else {
			playerType = 2
		}

		//checks turn
		if playerType != game.turn {
			return "Error: This is not your turn"
		}

		//changes turn
		if game.player1 == uid {
			games[gameId] = Game{
				player1: game.player1,
				player2: game.player2,
				turn:    1,
			}
		} else {
			games[gameId] = Game{
				player1: game.player1,
				player2: game.player2,
				turn:    2,
			}
		}

		s.Leave(gameId)
		io.BroadcastToRoom("/", gameId, "madeTurn", cords)
		s.Join(gameId)

		return ""
	})
}
