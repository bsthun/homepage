package file

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetDscFiles(dir string) ([]fs.DirEntry, error) {
	// * Get absolute file path
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to get absolute path: %s", err.Error())
	}

	// * Check that the folder exist
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory is not exist")
	}

	// * Read directory information
	files, err := os.ReadDir(abs)
	if err != nil {
		log.Fatal(err)
	}

	// * Filter only image files
	var filtered []fs.DirEntry
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") && strings.HasSuffix(strings.ToLower(file.Name()), ".jpg") {
			filtered = append(filtered, file)
		}
	}

	return filtered, nil
}
