package users

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Data folder not found creating..")
		//no folder found
		os.Mkdir("./data", 0755)
		err = writeFile([]byte("{}"))

		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = os.Stat("./data/users.json")
	if os.IsNotExist(err) {
		fmt.Println("users.json not found creating...")
		//only file doent not found
		err = writeFile([]byte("{}"))

		if err != nil {
			return err
		}
	} else if err != nil {
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

func GetData() (map[string]string, error) {
	bytes, err := Get()

	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func Set(id string, name string) error {
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

func ResetFile() error {
	checkForDataFile()
	return writeFile([]byte("{}"))
}
