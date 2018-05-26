package database

import (
	"fmt"
	"log"

	bolt "github.com/coreos/bbolt"
)

const TaskBucket = "Tasks"

func OpenDB() *bolt.DB {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatalf("Could not open DB: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TaskBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
