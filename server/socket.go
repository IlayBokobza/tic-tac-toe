package main

import (
	"encoding/json"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/ilaybokobza/tic-tac-toe/server/games"
	"github.com/ilaybokobza/tic-tac-toe/server/users"
)

func socketEvents(io *socketio.Server) {
	io.OnConnect("/", func(s socketio.Conn) error {
		fmt.Printf("\nnew connection \nid is %v", s.ID())
		return nil
	})

	io.OnEvent("/", "createGame", func(s socketio.Conn) string {
		id := createId()
		uid := s.ID()
		s.Join(id)
		users.Add(id, uid)
		games.Add(id, games.Game{Player1: uid})

		return id
	})

	io.OnEvent("/", "joinGame", func(s socketio.Conn, id string) string {
		gamesBytes, err := games.Get()

		if err != nil {
			return err.Error()
		}

		var gamesData map[string]games.Game
		json.Unmarshal(gamesBytes, &gamesData)

		//checks if game exists
		if len(gamesData[id].Player1) == 0 {
			return "Error: Game not found"
		}

		//check if game is full
		if len(gamesData[id].Player2) != 0 {
			return "Error: Game is full"
		}

		//joins him to game
		users.Add(id, s.ID())
		games.Add(id, games.Game{Player1: gamesData[id].Player1, Player2: s.ID(), Turn: 1})
		io.BroadcastToRoom("/", id, "startGame")
		s.Join(id)

		return ""
	})

	io.OnEvent("/", "madeTurn", func(s socketio.Conn, cords []int) string {
		uid := s.ID()
		usersBytes, err := users.Get()

		if err != nil {
			return err.Error()
		}

		gamesBytes, err := games.Get()

		if err != nil {
			return err.Error()
		}

		var usersData map[string]string
		var gamesData map[string]games.Game

		json.Unmarshal(usersBytes, &usersData)
		json.Unmarshal(gamesBytes, &gamesData)

		gameId := usersData[uid]
		game := gamesData[gameId]
		var playerType int

		fmt.Printf("\nuser id is %v the user map is: \n", uid)
		fmt.Println(usersData)

		if uid == game.Player1 {
			playerType = 1
		} else {
			playerType = 2
		}

		//checks turn
		if playerType != game.Turn {
			fmt.Printf("\n player type is %v but the turn is %v", playerType, game.Turn)
			return "Error: This is not your turn"
		}

		//changes turn
		if game.Player1 == uid {
			games.Add(gameId, games.Game{
				Player1: game.Player1,
				Player2: game.Player2,
				Turn:    1,
			})
		} else {
			games.Add(gameId, games.Game{
				Player1: game.Player1,
				Player2: game.Player2,
				Turn:    2,
			})
		}

		s.Leave(gameId)
		io.BroadcastToRoom("/", gameId, "madeTurn", cords)
		s.Join(gameId)

		return ""
	})
}
