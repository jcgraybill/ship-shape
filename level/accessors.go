package level

import (
	"fmt"

	"github.com/jcgraybill/ship-shape/planet"
)

func (lvl *Level) Planets() []*planet.Planet {
	return lvl.planets
}

func (lvl *Level) AllowedStructures() []int {
	return lvl.allowedStructures
}

func (lvl *Level) AllowedResources() []int {
	return lvl.allowedResources
}

func (lvl *Level) StartingMoney() uint {
	return lvl.startingMoney
}

func (lvl *Level) Message() string {
	return lvl.message
}

func (lvl *Level) Label() string {
	if lvl.progress >= lvl.goal {
		return fmt.Sprintf("%s (%d/%d): DONE", lvl.label, lvl.goal, lvl.goal)
	} else {
		return fmt.Sprintf("%s (%d/%d):", lvl.label, lvl.progress, lvl.goal)
	}
}

func (lvl *Level) ProgressBarValue() uint8 {
	return uint8(255 * float32(lvl.progress) / float32(lvl.goal))
}

func (lvl *Level) NextLevel() *Level {
	return lvl.nextLevel
}

func (lvl *Level) Title() string {
	return lvl.title
}
