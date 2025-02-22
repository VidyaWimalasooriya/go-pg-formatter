package format

import (
	"fmt"

	"os"
	"path/filepath"
)

func findFiles(path string) ([]string, error) {
	filePaths := []string{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			if ext == SQLFileExtension || ext == PostgresFileExtension {
				filePaths = append(filePaths, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", path, err)
	}

	return filePaths, nil
}
