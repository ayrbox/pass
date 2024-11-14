package utils

import (
	"os"

	path "github.com/muesli/go-app-paths"
)

func SetupPath(dirName string) (string, error) {
	scope := path.NewScope(path.User, dirName)
	dirs, err := scope.DataDirs()
	if err != nil {
		return "", err
	}

	// create the app base dir, if it doesn't exist
	var path string
	if len(dirs) > 0 {
		path = dirs[0]
	} else {
		path, _ = os.UserHomeDir()
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, 0o770)
			return path, nil
		}
		return "", err
	}
	return path, nil
}
