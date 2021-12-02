package main

func NewInMemoryScoreStore() *InMemoryScoreStore {
	return &InMemoryScoreStore{map[string]int{}}
}

type InMemoryScoreStore struct {
	store map[string]int
}

func (i *InMemoryScoreStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryScoreStore) RecordWin(name string) {
	i.store[name]++
}
