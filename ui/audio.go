package ui

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var buttonSound *audio.Player

func init() {
	audio.NewContext(44100)
	audioContext := audio.CurrentContext()
	audioBytes, err := GameData("audio/button.ogg")
	if err == nil {
		d, err := vorbis.Decode(audioContext, bytes.NewReader(audioBytes))
		if err == nil {
			buttonSound, err = audioContext.NewPlayer(d)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func PlayButtonSound() {
	buttonSound.Rewind()
	buttonSound.Play()
}
