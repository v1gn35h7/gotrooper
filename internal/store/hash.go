package store

import (
	"github.com/v1gn35h7/gotrooper/pb"
)

type TrooperStore struct {
	kvMap map[string]pb.ShellScript
}

func (t *TrooperStore) Set(key string, value pb.ShellScript) {
	t.kvMap[key] = value
}

func (t *TrooperStore) Get(key string) pb.ShellScript {

	if _, ok := t.kvMap[key]; ok == false {
		return pb.ShellScript{}
	}

	return t.kvMap[key]
}

func NewTrooperStore() *TrooperStore {
	return new(TrooperStore)
}
