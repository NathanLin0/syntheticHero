package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NathanLin0/syntheticHero/env"
)

// 建立所有移動變化
//
// @param canMoveToDiagonal 是否可以移動到斜對角線
func createAllMoveIndexesMap(canMoveedToDiagonal bool) map[string][]int {
	moveIndexesMap := map[string][]int{}
	for i := 0; i < env.Config.Row*env.Config.Column; i += env.Config.Row {
		// 橫向
		for j := 0; j < env.Config.Row-1; j++ {
			currentIndex := i + j
			moveIndex := currentIndex + 1
			moveIndexesMap[fmt.Sprintf("%d^%d", currentIndex, moveIndex)] = []int{currentIndex, moveIndex}
		}

		// 直向 左斜 右斜
		for j := 0; j < env.Config.Row; j++ {
			if env.Config.Row*env.Config.Column > i+j+env.Config.Row {
				currentIndex := i + j
				moveIndex := currentIndex + env.Config.Row
				moveIndexesMap[fmt.Sprintf("%d^%d", currentIndex, moveIndex)] = []int{currentIndex, moveIndex}

				if canMoveedToDiagonal {
					// 左下斜
					lowerLeftBevelIndex := (currentIndex + env.Config.Row - 1)
					if (i+env.Config.Row-1) < lowerLeftBevelIndex && (i+env.Config.Row+env.Config.Column) > lowerLeftBevelIndex {
						moveIndexesMap[fmt.Sprintf("%d^%d", currentIndex, lowerLeftBevelIndex)] = []int{currentIndex, lowerLeftBevelIndex}
					}

					// 右斜
					lowerRightBevel := (currentIndex + env.Config.Row + 1)
					if (i+env.Config.Row-1) < lowerRightBevel && (i+env.Config.Row+env.Config.Column) > lowerRightBevel {
						moveIndexesMap[fmt.Sprintf("%d^%d", currentIndex, lowerRightBevel)] = []int{currentIndex, lowerRightBevel}
					}
				}

			}
		}
	}
	return moveIndexesMap
}

func genMoveIndexesMap(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := createAllMoveIndexesMap(env.Config.CanMoveedToDiagonal)

	b, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		panic(marshalErr)
	}

	if _, writeErr := file.Write(b); err != nil {
		panic(writeErr)
	}
}
