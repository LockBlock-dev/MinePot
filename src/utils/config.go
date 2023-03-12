package utils

import (
	"encoding/json"
	"os"

	"github.com/LockBlock-dev/MinePot/typings"
)

func GetConfig() (*typings.Config, error) {
    file, err := os.Open("/etc/minepot/config.json")
    if err != nil {
        return &typings.Config{}, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := typings.Config{}
    err = decoder.Decode(&config)
    if err != nil {
        return &typings.Config{}, err
    }

    return &config, nil
}
