package storage

import (
	"devnotes/model"
	"sync"
)

type MemoryStore struct {
	Mu    sync.Mutex
	Users map[string]model.User
	Notes map[string]model.Note
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users: make(map[string]model.User),
		Notes: make(map[string]model.Note),
	}
}
