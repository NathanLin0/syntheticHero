package gamecore

import (
	"math/rand"

	"github.com/NathanLin0/syntheticHero/env"
)

func checkLayout(layout []int) bool {
	for i := range layout {
		if layout[i] == env.Config.AnyRoleTag {
			return true
		}
	}
	return false
}

func changeAnyRoleTag(layout []int) []int {
	if !checkLayout(layout) {
		return layout
	}

	for i := range layout {
		if layout[i] == env.Config.AnyRoleTag || layout[i] > 10 {
			role := env.Config.DefaultGenerateRoles[rand.Intn(len(env.Config.DefaultGenerateRoles))]
			layout[i] = role
		}
	}
	newLayout, _ := Start(layout, []int{}, -1)

	return changeAnyRoleTag(newLayout)
}

func CreateLayout() []int {
	layout := []int{}
	for i, l := 0, env.Config.Row*env.Config.Column; i < l; i++ {
		layout = append(layout, env.Config.AnyRoleTag)
	}

	return changeAnyRoleTag(layout)
}
