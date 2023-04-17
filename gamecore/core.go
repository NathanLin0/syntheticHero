package gamecore

import (
	"math"

	"github.com/NathanLin0/syntheticHero/env"
	"github.com/NathanLin0/syntheticHero/utils"
)

func getInitTagLayoutList(layoutLen int) []int {
	tagLayoutList := make([]int, layoutLen)
	for i := range tagLayoutList {
		tagLayoutList[i] = env.Config.AnyRoleTag
	}
	return tagLayoutList
}

func checkConnection(arr []int) map[int][]int {
	connectionMap := map[int][]int{}
	for i, r := range arr {
		if _, ok := connectionMap[r]; !ok {
			connectionMap[r] = []int{}
		}
		if len(connectionMap[r]) > 0 {
			if math.Abs(float64(connectionMap[r][len(connectionMap[r])-1]-i)) > 1 {
				if len(connectionMap[r]) >= env.Config.MinimumConnections {
					continue
				}
				connectionMap[r] = []int{i}
				continue
			}
		}
		connectionMap[r] = append(connectionMap[r], i)
	}
	for r, list := range connectionMap {
		if r == env.Config.ChestRole {
			continue
		}
		if len(list) < env.Config.MinimumConnections {
			delete(connectionMap, r)
		}
	}

	return connectionMap
}

func tagRoles(list, currentUpgradeList, moveIndexes []int, connectionMap map[int][]int) []int {
	newList := make([]int, len(list))
	copy(newList, currentUpgradeList)

	for r, indexes := range connectionMap {
		if _, ok := env.Config.UpgradeRolesMap[r]; !ok {
			continue
		}
		// 消除金幣
		if r == env.Config.CoinRole {
			upgradeIndex := -1
			for _, index := range moveIndexes {
				if idx := utils.IndexOf(indexes, index); idx != -1 {
					upgradeIndex = indexes[idx]
					break
				}
			}

			if upgradeIndex == -1 {
				if len(indexes) == 4 {
					upgradeIndex = indexes[1]
				} else {
					upgradeIndex = indexes[int(math.Floor(float64(len(indexes))/float64(2)))]
				}
			}

			for upgradeIndex > 0 && currentUpgradeList[upgradeIndex] == env.Config.UpgradeRoleTag {
				upgradeIndex -= 1
			}

			for upgradeIndex < len(currentUpgradeList) && currentUpgradeList[upgradeIndex] == env.Config.UpgradeRoleTag {
				upgradeIndex += 1
			}

			for _, index := range indexes {
				newList[index] = env.Config.EliminateRole
				if upgradeIndex == index {
					newList[index] = env.Config.UpgradeRoleTag
				}
			}
			continue
		}

		// 4消
		if len(indexes) == 4 {

			if currentUpgradeList[indexes[1]] == env.Config.UpgradeRoleTag {
				newList[indexes[0]] = env.Config.UpgradeRoleTag
				newList[indexes[1]] = env.Config.EliminateRole
				newList[indexes[2]] = env.Config.UpgradeRoleTag
				newList[indexes[3]] = env.Config.EliminateRole
				continue
			}

			if currentUpgradeList[indexes[2]] == env.Config.UpgradeRoleTag {
				newList[indexes[0]] = env.Config.UpgradeRoleTag
				newList[indexes[1]] = env.Config.UpgradeRoleTag
				newList[indexes[2]] = env.Config.EliminateRole
				newList[indexes[3]] = env.Config.EliminateRole
				continue
			}

			newList[indexes[0]] = env.Config.EliminateRole
			newList[indexes[1]] = env.Config.UpgradeRoleTag
			newList[indexes[2]] = env.Config.UpgradeRoleTag
			newList[indexes[3]] = env.Config.EliminateRole
			continue
		}

		// 5消
		if len(indexes) >= 5 {
			if currentUpgradeList[indexes[1]] == env.Config.UpgradeRoleTag {

				newList[indexes[0]] = env.Config.UpgradeRoleTag
				newList[indexes[1]] = env.Config.EliminateRole
				newList[indexes[2]] = env.Config.UpgradeRoleTag
				newList[indexes[3]] = env.Config.UpgradeRoleTag
				newList[indexes[4]] = env.Config.EliminateRole
				continue
			}
			if currentUpgradeList[indexes[2]] == env.Config.UpgradeRoleTag {

				newList[indexes[0]] = env.Config.UpgradeRoleTag
				newList[indexes[1]] = env.Config.UpgradeRoleTag
				newList[indexes[2]] = env.Config.EliminateRole
				newList[indexes[3]] = env.Config.UpgradeRoleTag
				newList[indexes[4]] = env.Config.EliminateRole
				continue
			}
			if currentUpgradeList[indexes[3]] == env.Config.UpgradeRoleTag {

				newList[indexes[0]] = env.Config.UpgradeRoleTag
				newList[indexes[1]] = env.Config.UpgradeRoleTag
				newList[indexes[2]] = env.Config.UpgradeRoleTag
				newList[indexes[3]] = env.Config.EliminateRole
				newList[indexes[4]] = env.Config.EliminateRole
				continue
			}
			newList[indexes[0]] = env.Config.EliminateRole
			newList[indexes[1]] = env.Config.UpgradeRoleTag
			newList[indexes[2]] = env.Config.UpgradeRoleTag
			newList[indexes[3]] = env.Config.UpgradeRoleTag
			newList[indexes[4]] = env.Config.EliminateRole
			continue
		}

		// 3消
		upgradeIndex := -1
		for _, index := range moveIndexes {
			if idx := utils.IndexOf(indexes, index); idx != -1 {
				upgradeIndex = indexes[idx]
				break
			}
		}

		if upgradeIndex == -1 {
			upgradeIndex = indexes[int(math.Floor(float64(len(indexes))/float64(2)))]
		}

		for upgradeIndex > 0 && currentUpgradeList[upgradeIndex] == env.Config.UpgradeRoleTag {
			upgradeIndex -= 1
		}

		for upgradeIndex < len(currentUpgradeList) && currentUpgradeList[upgradeIndex] == env.Config.UpgradeRoleTag {
			upgradeIndex += 1
		}

		for _, index := range indexes {
			newList[index] = env.Config.EliminateRole
			if upgradeIndex == index {
				newList[index] = env.Config.UpgradeRoleTag
			}
		}
		continue
	}
	return newList
}

func upgradeRoles(tagLayoutList, layout []int) {
	for i := range layout {
		if tagLayoutList[i] == env.Config.UpgradeRoleTag {
			layout[i] = env.Config.UpgradeRolesMap[layout[i]]
		}
	}
}

func removeEliminateRole(tagLayoutList, layout []int) {
	for i := 0; i < env.Config.Row; i++ {
		// 列 列表
		columnIndexes := []int{}
		for j := 0; j < env.Config.Column; j++ {
			columnIndexes = append(columnIndexes, i+j*env.Config.Row)
		}

		for j := 0; j < len(columnIndexes); j++ {
			for tagLayoutList[columnIndexes[j]] == env.Config.EliminateRole {
				isAllEliminateRole := true
				for k := j; k < len(columnIndexes); k++ {
					if tagLayoutList[columnIndexes[k]] != env.Config.EliminateRole {
						isAllEliminateRole = false
					}
				}
				if isAllEliminateRole {
					break
				}
				for k := j; k < len(columnIndexes); k++ {
					if len(columnIndexes) > k+1 {
						oldTag := tagLayoutList[columnIndexes[k]]
						tagLayoutList[columnIndexes[k]] = tagLayoutList[columnIndexes[k+1]]
						tagLayoutList[columnIndexes[k+1]] = oldTag

						oldRole := layout[columnIndexes[k]]
						layout[columnIndexes[k]] = layout[columnIndexes[k+1]]
						layout[columnIndexes[k+1]] = oldRole
					}
					if columnIndexes[k] != env.Config.EliminateRole {
						isAllEliminateRole = true
					}
				}
			}
		}
	}

	for i, tag := range tagLayoutList {
		if tag == env.Config.EliminateRole {
			tagLayoutList[i] = env.Config.AnyRoleTag
			layout[i] = env.Config.AnyRoleTag
		}
	}
}

func removeChestRole(layout []int) {
	for i := 0; i < env.Config.Row; i++ {
		// 列 列表
		columnIndexes := []int{}
		for j := 0; j < env.Config.Column; j++ {
			columnIndexes = append(columnIndexes, i+j*env.Config.Row)
		}

		for j := 0; j < len(columnIndexes); j++ {
			if layout[columnIndexes[j]] == env.Config.ChestRole {
				if len(columnIndexes) > j+1 && layout[columnIndexes[j+1]] != env.Config.ChestRole {
					oldRole := layout[columnIndexes[j]]
					layout[columnIndexes[j]] = layout[columnIndexes[j+1]]
					layout[columnIndexes[j+1]] = oldRole
				} else {
					layout[columnIndexes[j]] = env.Config.AnyRoleTag
				}
			}
		}
	}
}

func tidyLayout(layout, moveIndexes []int, times int) ([]int, int) {
	tagLayoutList := getInitTagLayoutList(len(layout))

	for i := 0; i < len(layout); i += env.Config.Row {
		// 行 列表
		rowList := layout[i : i+env.Config.Row]
		rowUpgradeList := tagLayoutList[i : i+env.Config.Row]
		rowConnectionMap := checkConnection(rowList)
		rowMoveIndexes := []int{}
		for _, index := range moveIndexes {
			if index >= i && index < i+env.Config.Row {
				rowMoveIndexes = append(rowMoveIndexes, index-i)
			}
		}
		rowTagResultList := tagRoles(rowList, rowUpgradeList, rowMoveIndexes, rowConnectionMap)

		for j, tag := range rowTagResultList {
			if tag != env.Config.AnyRoleTag && tagLayoutList[j+i] != env.Config.UpgradeRoleTag {
				tagLayoutList[j+i] = tag
			}
		}
	}

	for i := 0; i < env.Config.Row; i++ {
		// 列 列表
		columnList := []int{}
		rowUpgradeList := []int{}
		columnMoveIndexes := []int{}
		for j := 0; j < env.Config.Column; j++ {
			columnList = append(columnList, layout[i+j*env.Config.Row])
			rowUpgradeList = append(rowUpgradeList, tagLayoutList[i+j*env.Config.Row])
			resultIndex := utils.IndexOf(moveIndexes, i+j*env.Config.Row)
			if resultIndex >= 0 {
				columnMoveIndexes = append(columnMoveIndexes, j)
			}
		}
		columnConnectionMap := checkConnection(columnList)

		columnTagResultList := tagRoles(columnList, rowUpgradeList, columnMoveIndexes, columnConnectionMap)

		for j, tag := range columnTagResultList {
			if tag != env.Config.AnyRoleTag && tagLayoutList[i+j*env.Config.Row] != env.Config.UpgradeRoleTag {
				tagLayoutList[i+j*env.Config.Row] = tag
			}
		}
	}

	isAllTagAny := true
	for _, tag := range tagLayoutList {
		if tag != env.Config.AnyRoleTag {
			isAllTagAny = false
		}
	}
	if isAllTagAny {
		removeChestRole(layout)
		return layout, times
	}

	upgradeRoles(tagLayoutList, layout)

	removeEliminateRole(tagLayoutList, layout)

	times++
	return tidyLayout(layout, []int{}, times)
}

func Start(layout []int, moveIndexes []int, removeIndex int) ([]int, int) {
	if len(layout) != env.Config.Row*env.Config.Column {
		panic("layout size filed")
	}
	newLayout := make([]int, len(layout))
	copy(newLayout, layout)

	// 刪除固定位置
	if removeIndex >= 0 && removeIndex < len(newLayout) {
		tagLayoutList := getInitTagLayoutList(len(newLayout))
		tagLayoutList[removeIndex] = env.Config.EliminateRole
		removeEliminateRole(tagLayoutList, newLayout)
	}

	// 移動位置 互換
	if len(moveIndexes) == 2 {
		oldRole := newLayout[moveIndexes[0]]
		newLayout[moveIndexes[0]] = newLayout[moveIndexes[1]]
		newLayout[moveIndexes[1]] = oldRole
	}

	resultLayout, times := tidyLayout(newLayout, moveIndexes, 0)

	return resultLayout, times
}
