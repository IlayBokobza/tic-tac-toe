package users

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func checkForDataFile() error {
	//checks for folder
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		//no folder found
		os.Mkdir("./data", 0755)
		err = ioutil.WriteFile("./data/users.json", []byte("[]"), 0666)
		return err
	}

	_, err = os.Stat("./data/users.json")
	if os.IsNotExist(err) {
		//only file doent not found
		err = ioutil.WriteFile("./data/users.json", []byte("[]"), 0666)
		return err
	}

	return nil
}

func Get() ([]byte, error) {
	//creates file/folder if does not exsit
	checkForDataFile()
	//gets data
	return ioutil.ReadFile("./data/users.json")
}

func AddUser(id string, name string) error {
	file, err := Get()

	if err != nil {
		return err
	}

	var data []map[string]string
	err = json.Unmarshal(file, &data)

	if err != nil {
		return err
	}

	data = append(data, map[string]string{name: id})
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./data/users.json", json, 0666)

	if err != nil {
		return err
	}

	return nil
}

func Delete(name string) error {
	data, err := Get()

	if err != nil {
		return err
	}

	var users map[string]string
	err = json.Unmarshal(data, &users)

	if err != nil {
		return err
	}

	return nil
}
