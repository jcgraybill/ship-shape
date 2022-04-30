package level

import (
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

func (lvl *Level) ShowStatus() (string, string, uint, uint) {
	if lvl.goalMet {
		return lvl.message, lvl.label, lvl.goal, lvl.goal
	}
	return lvl.message, lvl.label, lvl.progress, lvl.goal
}

func (lvl *Level) NextLevel() *Level {
	return lvl.nextLevel
}

func (lvl *Level) Title() string {
	return lvl.title
}
