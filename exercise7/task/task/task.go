package task

import (
	"fmt"
	"gophercise/exercise7/task/database"

	bolt "github.com/coreos/bbolt"
)

const (
	complete   = "Complete"
	incomplete = "Incomplete"
)

// Add adds a task to the provided database
func Add(taskName string, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.TaskBucket))
		err := b.Put([]byte(taskName), []byte(incomplete))
		return err
	})
}

// Do updates the value of the specified task to complete
func Do(taskName string, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.TaskBucket))
		err := b.Put([]byte(taskName), []byte(complete))
		return err
	})
}

// List prints all of the incomplete tasks
func List(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.TaskBucket))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) == incomplete {
				fmt.Println(string(k))
			}
		}


		return nil
	})
}
