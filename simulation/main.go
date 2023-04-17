package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/NathanLin0/syntheticHero/env"
	"github.com/NathanLin0/syntheticHero/gamecore"
)

var wg sync.WaitGroup
var w sync.Mutex

type Result struct {
	NextIndex string `json:"nextIndex"`
	Times     int    `json:"times"`
}

var data map[string]bool

func init() {
	rand.Seed(time.Now().Unix())
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data = map[string]bool{}

	simulationFilePath := path.Join(dir, "..", "conf.d", "simulation.log")
	file, err := os.Open(simulationFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(simulationFilePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	defer file.Close()

	// 创建带缓冲的读取器
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		result := map[string]Result{}
		if unmarshalErr := json.Unmarshal([]byte(line), &result); unmarshalErr != nil {
			panic(unmarshalErr)
		}
		for k := range result {
			data[k] = true
		}
	}
}

func formatKey(layout []int, moveIndexesKey, removeIndexKey string) string {
	layoutString := []string{}
	for _, val := range layout {
		layoutString = append(layoutString, fmt.Sprint(val))
	}
	return fmt.Sprintf("%s.%s.%s", strings.Join(layoutString, "^"), moveIndexesKey, removeIndexKey)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	env.Init(path.Join(dir, ".."))

	simulationFilePath := path.Join(dir, "..", "conf.d", "simulation.log")
	file, err := os.OpenFile(simulationFilePath, os.O_WRONLY|os.O_APPEND, 06444)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(simulationFilePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	defer file.Close()

	for i := 0; i < 1000; i++ {
		layout := gamecore.CreateLayout()
		wg.Add(1)
		go run(layout, file)
	}

	wg.Wait()
}

func search(key string) bool {
	w.Lock()
	defer w.Unlock()
	_, ok := data[key]
	return ok
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// simulationFilePath := path.Join(dir, "..", "conf.d", "simulation.log")
	// file, err := os.Open(simulationFilePath)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		file, err = os.Create(simulationFilePath)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	} else {
	// 		panic(err)
	// 	}
	// }
	// defer file.Close()

	// // 创建带缓冲的读取器
	// reader := bufio.NewReader(file)

	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		return false
	// 	}
	// 	result := map[string]Result{}
	// 	if unmarshalErr := json.Unmarshal([]byte(line), &result); unmarshalErr != nil {
	// 		panic(unmarshalErr)
	// 	}
	// 	if _, ok := result[key]; ok {
	// 		return true
	// 	}
	// }
}

func run(layout []int, file *os.File) {
	checkMap := map[int]int{}
	for _, v := range layout {
		checkMap[v]++
	}

	defer wg.Done()

	// isRun := false
	// for role, count := range checkMap {
	// 	if role != env.Config.AnyRoleTag && count >= env.Config.MinimumConnections {
	// 		isRun = true
	// 	}
	// }

	if checkMap[env.Config.AnyRoleTag] > 0 {
		return
	}

	// for k, removeIndex := range env.RemoveIndexes {
	// 	key := formatKey(layout, "", k)
	// 	if !search(key) {
	// 		resultLayout, times := gamecore.Start(layout, []int{}, removeIndex)
	// 		appendFile(key, formatKey(resultLayout, "", k), times, file)
	//      wg.Add(1)
	// 		go run(resultLayout, file)
	// 	}
	// }

	for k, moveIndex := range env.MoveIndexes {
		key := formatKey(layout, k, "")
		if !search(key) && layout[moveIndex[0]] != layout[moveIndex[1]] {
			resultLayout, times := gamecore.Start(layout, moveIndex, -1)
			appendFile(key, formatKey(resultLayout, k, ""), times, file)
			wg.Add(1)
			go run(resultLayout, file)
		}
	}
}

func appendFile(key, nextKey string, times int, file *os.File) {
	w.Lock()

	defer w.Unlock()
	b, err := json.Marshal(map[string]Result{
		key: {
			NextIndex: nextKey,
			Times:     times,
		},
	})
	if err != nil {
		panic(err)
	}

	data[key] = true
	b = append(b, []byte("\n")...)
	file.Write(b)
	file.Sync()
}
