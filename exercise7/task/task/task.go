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
func List(db *bolt.DB) []string {
	var ret []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.TaskBucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) == incomplete {
				ret = append(ret, string(k))
			}
		}

		return nil
	})

	return ret
}

func PrintTasks(tasks []string) {
	switch len(tasks) {
	case 0:
		fmt.Printf("No incomplete tasks found\nNice work!\n")
	default:
		fmt.Printf("Here are your current tasks:\n")
		for i, t := range tasks {
			fmt.Printf("#%v : %v\n", i+1, t)
		}
	}

}
