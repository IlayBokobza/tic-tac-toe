package main

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/ilaybokobza/tic-tac-toe/server/games"
	"github.com/ilaybokobza/tic-tac-toe/server/users"
)

func socketEvents(io *socketio.Server) {
	io.OnConnect("/", func(s socketio.Conn) error {
		fmt.Printf("New Connection to server with the id of \"%v\". \n", s.ID())
		return nil
	})

	//user creates game
	io.OnEvent("/", "createGame", func(s socketio.Conn) string {
		uid := s.ID()
		//checks if user has create a game aleady
		usersData, _ := users.GetData()
		if len(usersData[uid]) != 0 {
			return usersData[uid]
		}

		id := games.CreateID()
		s.Join(id)
		users.Set(id, uid)
		games.Set(id, games.Game{Player1: uid})

		return id
	})

	//user joins game
	io.OnEvent("/", "joinGame", func(s socketio.Conn, id string) string {
		gamesData, err := games.GetData()

		if err != nil {
			return err.Error()
		}

		//checks if game exists
		if len(gamesData[id].Player1) == 0 {
			return "Error: Game not found"
		}

		//check if game is full
		if len(gamesData[id].Player2) != 0 {
			return "Error: Game is full"
		}

		//joins him to game
		users.Set(id, s.ID())
		games.Set(id, games.Game{Player1: gamesData[id].Player1, Player2: s.ID(), Turn: 1, Board: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}})
		io.BroadcastToRoom("/", id, "startGame")
		s.Join(id)

		return ""
	})

	//user asks if they are in game
	io.OnEvent("/", "isInGame", func(s socketio.Conn) string {
		usersData, err := users.GetData()

		if err != nil {
			return err.Error()
		}

		if len(usersData[s.ID()]) == 0 {
			return "Error: you are not in a game"
		}

		return ""
	})

	//user makes a turn
	io.OnEvent("/", "madeTurn", func(s socketio.Conn, cords []int) string {
		usersData, err := users.GetData()

		if err != nil {
			return err.Error()
		}

		gamesData, err := games.GetData()

		if err != nil {
			return err.Error()
		}

		uid := s.ID()
		gameId := usersData[uid]
		game := gamesData[gameId]
		var playerType int

		if uid == game.Player1 {
			playerType = 1
		} else {
			playerType = 2
		}

		//checks turn
		if playerType != game.Turn {
			return "Error: This is not your turn."
		}

		//checks thats spot isnt taken
		y := cords[0]
		x := cords[1]

		if game.Board[x][y] > 0 {
			return "Error: Spot is taken."
		}

		//changes board
		game.Board[x][y] = playerType

		//checks for win
		if game.TurnsMade+1 >= 5 {
			winState := games.CheckForWin(game.Board)

			if winState != "none" {
				s.Leave(gameId)
				io.BroadcastToRoom("/", gameId, "madeTurn", cords)
				s.Join(gameId)
				io.BroadcastToRoom("/", gameId, "win", winState)
				return ""
			}
		}

		//checks for tie
		if game.TurnsMade+1 == 9 {
			s.Leave(gameId)
			io.BroadcastToRoom("/", gameId, "madeTurn", cords)
			s.Join(gameId)
			io.BroadcastToRoom("/", gameId, "tie")
			return ""
		}

		//swaps turns
		var newTurn int
		if playerType == 1 {
			newTurn = 2
		} else {
			newTurn = 1
		}

		//updates data
		games.Set(gameId, games.Game{
			Player1:   game.Player1,
			Player2:   game.Player2,
			Turn:      newTurn,
			Board:     game.Board,
			TurnsMade: game.TurnsMade + 1,
		})

		s.Leave(gameId)
		io.BroadcastToRoom("/", gameId, "madeTurn", cords)
		s.Join(gameId)

		return ""
	})

	io.OnEvent("/", "askForNewGame", func(s socketio.Conn) {
		usersData, err := users.GetData()

		if err != nil {
			return
		}

		gameId := usersData[s.ID()]

		s.Leave(gameId)
		io.BroadcastToRoom("/", gameId, "askForNewGame")
		s.Join(gameId)
	})

	io.OnEvent("/", "newGame", func(s socketio.Conn) {
		usersData, err := users.GetData()

		if err != nil {
			return
		}

		gameId := usersData[s.ID()]

		err = games.Reset(gameId)

		if err != nil {
			return
		}

		io.BroadcastToRoom("/", gameId, "newGame")
	})

	io.OnEvent("/", "endGame", func(s socketio.Conn) {
		clearGame(io, s.ID())
	})

	io.OnDisconnect("/", func(s socketio.Conn, _ string) {
		clearGame(io, s.ID())
	})
}

//clears game by one of the players
func clearGame(io *socketio.Server, uid string) {
	usersData, err := users.GetData()

	if err != nil {
		return
	}

	gamesData, err := games.GetData()

	if err != nil {
		return
	}

	gameId := usersData[uid]
	game := gamesData[gameId]

	//if user wasnt on the list
	if len(gameId) == 0 {
		return
	}

	//deletes him and his game
	users.Delete(game.Player1)
	users.Delete(game.Player2)
	games.Delete(gameId)

	//clear room
	io.BroadcastToRoom("/", gameId, "gameOver")
	io.ClearRoom("/", gameId)
}
