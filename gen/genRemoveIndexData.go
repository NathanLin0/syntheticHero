package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NathanLin0/syntheticHero/env"
)

func createAllRemoveIndexMap() map[string]int {
	removeIndexMap := map[string]int{}
	for i := 0; i < env.Config.Row*env.Config.Column; i++ {
		removeIndexMap[fmt.Sprint(i)] = i
	}
	return removeIndexMap
}

func genRemoveIndexJsonFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := createAllRemoveIndexMap()

	b, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		panic(marshalErr)
	}

	if _, writeErr := file.Write(b); err != nil {
		panic(writeErr)
	}
}
