package env

import (
	"encoding/json"
	"io/ioutil"
	"path"

	_ "github.com/NathanLin0/syntheticHero/conf.d"
	"gopkg.in/yaml.v2"
)

var Config gameConfig
var RemoveIndexes map[string]int
var MoveIndexes map[string][]int

type generateConfig struct {
	CanMoveedToDiagonal bool `yaml:"canMoveedToDiagonal"`
}

type gameConfig struct {
	Row                  int         `yaml:"row"`
	Column               int         `yaml:"column"`
	MinimumConnections   int         `yaml:"minimumConnections"`
	MaximumConnections   int         `yaml:"maximumConnections"`
	EliminateRole        int         `yaml:"eliminateRole"`
	UpgradeRoleTag       int         `yaml:"upgradeRoleTag"`
	AnyRoleTag           int         `yaml:"anyRoleTag"`
	CoinRole             int         `yaml:"coinRole"`
	ChestRole            int         `yaml:"chestRole"`
	DefaultGenerateRoles []int       `yaml:"defaultGenerateRoles"`
	UpgradeRolesMap      map[int]int `yaml:"upgradeRolesMap"`
	generateConfig
}

func LoadYamlConfig[T any](path string, result *T) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if unmarshalErr := yaml.Unmarshal(data, result); unmarshalErr != nil {
		return unmarshalErr
	}
	return nil
}

func LoadJsonConfig[T any](path string, result *T) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if unmarshalErr := json.Unmarshal(data, result); unmarshalErr != nil {
		return unmarshalErr
	}
	return nil
}

func Init(dir string) {
	Config = gameConfig{}
	MoveIndexes = make(map[string][]int)
	RemoveIndexes = make(map[string]int)

	if err := LoadYamlConfig(path.Join(dir, "conf.d", "game.yaml"), &Config); err != nil {
		panic(err)
	}

	if err := LoadJsonConfig(path.Join(dir, "conf.d", "moveIndexes.json"), &MoveIndexes); err != nil {
		panic(err)
	}

	if err := LoadJsonConfig(path.Join(dir, "conf.d", "removeIndexes.json"), &RemoveIndexes); err != nil {
		panic(err)
	}
}
