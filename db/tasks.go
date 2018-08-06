package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task represent a TODO
type Task struct {
	Key   int
	Value string
}

// Init : Initializes the database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask : Creates a new task with new ID
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := intToByte(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetAllTasks : Retrieves all the tasks from the database
func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   byteToInt(k),
				Value: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// DeleteTask : Deletes a task
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(intToByte(key))
	})

}

// IsTaskListEmpty :returns true if task list empty
func IsTaskListEmpty() bool {
	var isEmpty = false
	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		lastkey, _ := c.Last()
		if lastkey == nil {
			isEmpty = true
		}
		return nil
	})
	return isEmpty
}

// ResetTaskSequence : Resets the task sequence back to 0
func ResetTaskSequence() error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.SetSequence(0)
	})
}

func intToByte(id int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b
}

func byteToInt(bytes []byte) int {
	return int(binary.BigEndian.Uint64(bytes))
}
