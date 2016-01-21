package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

type Config struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Environment string `json:"env"`
}

func ConfigReader(filepath string) (Config, error) {
	var config Config
	cfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Unable to read/find file", err)
		return config, err
	}
	err = json.Unmarshal(cfile, &config)
	if err != nil {
		fmt.Println("Unmarshal Error:", err)
	}
	return config, nil
}

func WriteToFile(data string, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func WriteBytesToFile(data []byte, file string) {
	ioutil.WriteFile(file, data, 777)
}

func ReadBytesFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func ReadFromFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data), err
}

func GetHomeDir() (homeDir string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
