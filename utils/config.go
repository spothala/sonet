package utils

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Config struct {
    Version     string  `json:"version"`
    Description string  `json:"description"`
    Environment string  `json:"env"`
}

func ConfigReader(filepath string)(Config, error) {
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
