package util

import (
	"encoding/json"
	"os"

	"github.com/LockBlock-dev/MinePot/types"
)

func GetConfig() (*types.Config, error) {
	file, err := os.Open("/etc/minepot/config.json")
	if err != nil {
		return &types.Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := types.Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return &types.Config{}, err
	}

	return &config, nil
}
