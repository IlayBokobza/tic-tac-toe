package users

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func writeFile(data []byte) error {
	return ioutil.WriteFile("./data/users.json", data, 0666)
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

	_, err = os.Stat("./data/users.json")
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
	return ioutil.ReadFile("./data/users.json")
}

func Add(id string, name string) error {
	file, err := Get()

	if err != nil {
		return err
	}

	var data map[string]string
	err = json.Unmarshal(file, &data)

	if err != nil {
		return err
	}

	data[name] = id
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

	delete(users, name)

	json, err := json.Marshal(users)

	if err != nil {
		return err
	}

	err = writeFile(json)

	if err != nil {
		return err
	}

	return nil
}
