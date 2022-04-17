//go:build deploy
// +build deploy

package ui

import (
	"embed"
)

func GameData(path string) ([]byte, error) {
	return gd.ReadFile(path)
}

//go:embed fonts audio
var gd embed.FS