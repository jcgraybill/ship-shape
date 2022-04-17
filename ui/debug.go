//go:build !deploy
// +build !deploy

package ui

import "os"

func GameData(path string) ([]byte, error) {
	return os.ReadFile("ui/" + path)
}
