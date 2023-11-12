package helpers

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

/**
 * ExpandPath(string) string
 *
 * Function receives a path (string) and trys to resolve/expand the XDG_HOME path.
 * If there is nothing to expand, we simply return the provided path.
 */
func ExpandPath(path string) string {
	usr, err := user.Current()
	if err != nil {
		return path
	}

	dir := usr.HomeDir
	if path == "~" || path == "$HOME" {
		path = dir
	} else if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "$HOME/") {
		path = JoinPath(dir, path[2:])
	}
	return path
}

/**
 * JoinPath(...string) string
 *
 * Returns the joined interface as a path.
 */
func JoinPath(pathParts ...string) string {
	return filepath.Join(pathParts...)
}

/**
 * CreateDirectory(string) error
 *
 * Creates the provided directory structure.
 * Returns nil or the thrown error.
 */
func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

/**
 * PathExsts(string) bool
 *
 * Retrun true if the path exists, false if not.
 */
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
 * GetAbsolutePath(string) string
 *
 * Returns the absolute path.
 */
func GetAbsolutePath(path string) string {
	path = ExpandPath(path)
	if !PathExists(path) {
		return path
	}
	// check if the provided path is a file or not
	if !FileExists(path) {
		return path
	}

	// remove the last element from the path becasue it's a file
	pathSplit := strings.Split(path, "/")
	return strings.Join(pathSplit[:len(pathSplit)-1], "/")
}
