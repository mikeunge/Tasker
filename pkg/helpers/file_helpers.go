package helpers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/**
 * WriteFile(path string, data string) error
 *
 * Write data to provided file path.
 */
func WriteFile(path string, data string) error {
	return os.WriteFile(path, []byte(data), 0644)
}

/**
 * CreateFile(path string) error
 *
 * Create a empty file, return error if action was not successful.
 */
func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

/**
 * ReadFileBytes(path string) ([]byte, error)
 *
 * Reads the provided file and returns the data as a byte array.
 */
func ReadFileBytes(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read data from file %s, %+v", path, err)
	}

	return data, nil
}

/**
 * ReadFile(string) (string, error)
 *
 * Reads the provided file and returns the data (string).
 */
func ReadFile(path string) (string, error) {
	data, err := ReadFileBytes(path)
	if err != nil {
		return "", err
	}

	return string(data[:]), nil
}

/**
 * FileExists(string) bool
 *
 * Returns true if the file exists, false if not.
 * False could also mean that the path exists but it's not a file.
 */
func FileExists(path string) bool {
	if info, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

/**
 * GetFileName(string) string
 *
 * Returns the file name from the provided path string.
 */
func GetFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
}

/**
 * GetFilesInPath(string) ([]string, error)
 *
 * Returns all the file in a given path + sub directories.
 */
func GetFilesInDir(path string) ([]string, error) {
	var files []string

	if !PathExists(path) {
		return files, fmt.Errorf("path '%s' does not exist", path)
	}

	// recursivly search for files in provided path
	err := filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return files, err
	}

	return files, nil
}
