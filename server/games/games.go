package games

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Player1   string  `json:"player1"`
	Player2   string  `json:"player2"`
	Turn      int     `json:"turn"`
	Board     [][]int `json:"board"`
	TurnsMade int     `json:"turnsMade"`
}

func writeFile(data []byte) error {
	return ioutil.WriteFile("./data/games.json", data, 0666)
}

func checkForDataFile() error {
	//checks for folder
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		//no folder found
		os.Mkdir("./data", 0755)
		err = writeFile([]byte("{}"))
		return err
	}

	_, err = os.Stat("./data/games.json")
	if os.IsNotExist(err) {
		//only file doent not found
		err = writeFile([]byte("{}"))
		return err
	}

	return nil
}

func Get() ([]byte, error) {
	//creates file/folder if does not exsit
	checkForDataFile()
	//gets data
	return ioutil.ReadFile("./data/games.json")
}

func GetData() (map[string]Game, error) {
	bytes, err := Get()

	if err != nil {
		return nil, err
	}

	var data map[string]Game
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func Set(id string, game Game) error {
	file, err := Get()

	if err != nil {
		return err
	}

	var data map[string]Game
	err = json.Unmarshal(file, &data)

	if err != nil {
		return err
	}

	data[id] = game
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = writeFile(json)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	file, err := Get()

	if err != nil {
		return err
	}

	var data map[string]Game
	err = json.Unmarshal(file, &data)

	if err != nil {
		return err
	}

	delete(data, id)

	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = writeFile(json)

	if err != nil {
		return err
	}

	return nil
}

func ResetFile() error {
	checkForDataFile()
	return writeFile([]byte("{}"))
}

func Reset(id string) error {
	file, err := Get()

	if err != nil {
		return err
	}

	var data map[string]Game
	err = json.Unmarshal(file, &data)

	if err != nil {
		return err
	}

	game := data[id]
	data[id] = Game{Player1: game.Player1, Player2: game.Player2, Turn: 1, TurnsMade: 0, Board: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}

	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = writeFile(json)

	if err != nil {
		return err
	}

	return nil
}

func CreateID() string {
	rand.Seed(time.Now().UnixNano())
	var id []byte

	for i := 0; i < 6; i++ {
		randInt := rand.Intn(122-97) + 97
		id = append(id, byte(randInt))
	}

	return string(id)
}

func checkAllPosabilities(board [][]int, p int) [][]int {
	//checks horizontal
	if board[0][0] == p && board[0][1] == p && board[0][2] == p {
		return [][]int{{0, 0}, {0, 1}, {0, 2}}
	}

	if board[1][0] == p && board[1][1] == p && board[1][2] == p {
		return [][]int{{1, 0}, {1, 1}, {1, 2}}
	}

	if board[2][0] == p && board[2][1] == p && board[2][2] == p {
		return [][]int{{2, 0}, {2, 1}, {2, 2}}
	}

	//checks vertical
	if board[0][0] == p && board[1][0] == p && board[2][0] == p {
		return [][]int{{0, 0}, {1, 0}, {2, 0}}
	}

	if board[0][1] == p && board[1][1] == p && board[2][1] == p {
		return [][]int{{0, 1}, {1, 1}, {2, 1}}
	}

	if board[0][2] == p && board[1][2] == p && board[2][2] == p {
		return [][]int{{0, 2}, {1, 2}, {2, 2}}
	}

	//checks diagonal
	if board[2][0] == p && board[1][1] == p && board[0][2] == p {
		return [][]int{{2, 0}, {1, 1}, {0, 2}}
	}

	if board[0][0] == p && board[1][1] == p && board[2][2] == p {
		return [][]int{{0, 0}, {1, 1}, {2, 2}}
	}

	return [][]int{}
}

func turnToJson(winPos [][]int) string {
	jsonBytes, _ := json.Marshal(winPos)

	return string(jsonBytes)
}

//checks a board for a win
func CheckForWin(board [][]int) string {
	//checks player 1
	winPos := checkAllPosabilities(board, 1)
	if len(winPos) > 0 {
		return turnToJson(winPos)
	}

	//checks player 2
	winPos = checkAllPosabilities(board, 2)
	if len(winPos) > 0 {
		return turnToJson(winPos)
	}

	return "none"
}
