package utils

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/LockBlock-dev/MinePot/typings"
)

func GetFavicon(config *typings.Config) (error) {
	faviconFile, err := os.ReadFile(config.FaviconPath)
    if err != nil {
        return fmt.Errorf("error reading the favicon file: %w", err)
    }
	
	config.StatusResponseData.Favicon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(faviconFile)

	return nil
}
