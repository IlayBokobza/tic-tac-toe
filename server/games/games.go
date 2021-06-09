package games

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Game struct {
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
	Turn    int    `json:"turn"`
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

func Add(id string, game Game) error {
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
