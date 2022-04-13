//go:build deploy
// +build deploy

package util

import (
	"embed"
)

func GameData(path string) ([]byte, error) {
	return gd.ReadFile(path)
}

//go:embed fonts
var gd embed.FS
