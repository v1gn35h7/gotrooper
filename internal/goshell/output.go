package goshell

import (
	"os"
	"sync"
)

type OutputFile struct {
	File  *os.File
	Mutex sync.RWMutex
}

type ScriptOutput struct {
	AgentId  string `json:"AgentId,omitempty"`
	HostName string `json:"Platform,omitempty"`
	ScriptId string `json:"ScriptId,omitempty"`
	Output   string `json:"Output,omitempty"`
}
