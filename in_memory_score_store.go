package main

import "sync"

func NewInMemoryScoreStore() *InMemoryScoreStore {
	return &InMemoryScoreStore{sync.RWMutex{}, map[string]int{}}
}

type InMemoryScoreStore struct {
	lock  sync.RWMutex
	store map[string]int
}

func (i *InMemoryScoreStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

func (i *InMemoryScoreStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}
