package main

import (
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/NathanLin0/syntheticHero/env"
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
	env.Init(path.Join(dir, ".."))

	genRemoveIndexJsonFile(path.Join(dir, "..", "conf.d", "removeIndexes.json"))

	genMoveIndexesMap(path.Join(dir, "..", "conf.d", "moveIndexes.json"))

	// fmt.Println(env.RemoveIndexes)
	// fmt.Println(env.MoveIndexes)

}
