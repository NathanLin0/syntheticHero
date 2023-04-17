package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/NathanLin0/syntheticHero/env"
	"github.com/NathanLin0/syntheticHero/gamecore"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Statistics struct {
	Layout      []int `json:"layout"`
	MoveIndexes []int `json:"moveIndexes"`
	Times       int   `json:"times"`
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	env.Init(dir)

	// layout := []int{
	// 	1, 2, 1, 2, 1, 2,
	// 	2, 1, 2, 1, 2, 1,
	// 	1, 2, 1, 2, 1, 2,
	// 	2, 1, 2, 1, 2, 1,
	// 	1, 2, 1, 2, 1, 2,
	// 	2, 1, 2, 1, 2, 1,
	// }

	// layout := []int{
	// 	5, 5, 1, 1, 4, 3,
	// 	1, 2, 1, 5, 3, 1,
	// 	4, 3, 1, 3, 4, 2,
	// 	2, 2, 1, 2, 4, 3,
	// 	4, 4, 4, 1, 1, 2,
	// 	2, 3, 1, 4, 4, 4,
	// }
	layout := gamecore.CreateLayout()
	// {"1^4^1^5^1^5^4^4^1^4^4^2^5^3^5^4^3^4^1^4^2^5^1^1^3^1^4^1^4^4^2^3^1^5^1^3..28":{"nextIndex":"1^4^1^5^1^5^4^4^1^4^4^2^5^3^5^4^3^4^1^4^2^5^1^1^3^1^4^1^1^4^2^3^1^5^-999^3..28","times":0}}

	// layout := []int{
	// 	-999, -999, -999, -999, -999, -999,
	// 	1, -999, 1, -999, -999, -999,
	// 	-999, 1, -999, -999, -999, -999,
	// 	-999, -999, -999, -999, -999, -999,
	// 	-999, -999, -999, -999, -999, -999,
	// 	-999, -999, -999, -999, -999, -999,
	// }

	// [1 2 10 2 1 2]
	// [2 10 2 2 2 1]
	// [1 2 1 1 1 2]
	// [2 1 2 2 2 1]
	// [1 2 -999 1 1 2]
	// [2 1 -999 -999 2 1]

	// fmt.Println("layout: ", layout)

	// for i := 0; i < len(layout); i += env.Config.Row {
	// 	// 行 列表
	// 	rowList := layout[i : i+env.Config.Row]
	// 	fmt.Println(rowList)
	// }

	// fmt.Println("------------------------------")

	newLayout, times := gamecore.Start(layout, []int{}, -1)
	fmt.Println("times: ", times)

	for i := 0; i < len(newLayout); i += env.Config.Row {
		// 行 列表
		rowList := newLayout[i : i+env.Config.Row]
		fmt.Println(rowList)
	}

	// for l := 0; l < len(layout); l += env.Config.Row {
	// 	for i := 0; i < env.Config.Row; i++ {
	// 		for j := i + 1; j < env.Config.Column; j++ {
	// 			if math.Abs(float64((l+i)-(l+j))) == 1 {
	// 				moveIndexes := []int{l + i, l + j}

	// 				times := gamecore.Start(layout, moveIndexes, -1)
	// 				fmt.Println("moveIndexes: ", moveIndexes)
	// 				fmt.Println("times: ", times)

	// 				for i := 0; i < len(layout); i += env.Config.Row {
	// 					// 行 列表
	// 					rowList := layout[i : i+env.Config.Row]
	// 					fmt.Println(rowList)
	// 				}

	// 				fmt.Println("------------------------------")

	// 			}

	// 			if math.Abs(float64((l+i)-(l+i+j*env.Config.Row))) == 6 {
	// 				moveIndexes2 := []int{l + i, l + i + j*env.Config.Row}

	// 				times2 := gamecore.Start(layout, moveIndexes2, -1)
	// 				fmt.Println("moveIndexes: ", moveIndexes2)
	// 				fmt.Println("times: ", times2)

	// 				for i := 0; i < len(layout); i += env.Config.Row {
	// 					// 行 列表
	// 					rowList := layout[i : i+env.Config.Row]
	// 					fmt.Println(rowList)
	// 				}
	// 				fmt.Println("------------------------------")
	// 			}
	// 		}
	// 	}
	// }
}
