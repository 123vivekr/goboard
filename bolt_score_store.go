package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

const ScoreStoreBucketName = "ScoreStore"

func NewBoltScoreStore() *BoltScoreStore {
	db, err := bolt.Open("goboard.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// TODO: close db when BoltScoreStore goes out of memory
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(ScoreStoreBucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &BoltScoreStore{sync.RWMutex{}, db}
}

type BoltScoreStore struct {
	mu sync.RWMutex
	db *bolt.DB
}

func (b *BoltScoreStore) GetPlayerScore(name string) int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	score := 0

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ScoreStoreBucketName))
		score, _ = strconv.Atoi(string(bucket.Get([]byte(name))))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return score
}

func (b *BoltScoreStore) RecordWin(name string) {
	prev_score_string := b.GetPlayerScore(name)

	b.mu.Lock()
	defer b.mu.Unlock()
	err := b.db.Update(func(tx *bolt.Tx) error {
		prev_score, _ := strconv.Atoi(fmt.Sprint(prev_score_string))

		bucket := tx.Bucket([]byte(ScoreStoreBucketName))
		err := bucket.Put([]byte(name), []byte(fmt.Sprintf("%d", prev_score+1)))
		if err != nil {
			return fmt.Errorf("record win: %s", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return
}
