package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"time"
)

var taskBucketName = []byte("tasks") //Table name is tasks
var db *bolt.DB

type Task struct {
	Key int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	createBucketFn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucketName)
		return err
	}
	
	return db.Update(createBucketFn)
}

func CreateTask(task string) (int, error) {
	var taskId int

	err := db.Update(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket(taskBucketName)
		id64, _ := taskBucket.NextSequence()
		taskId = int(id64)
		key := itob(taskId)
		return taskBucket.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return taskId, nil
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket(taskBucketName)

		c := taskBucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
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

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		taskBucket := tx.Bucket(taskBucketName)
		return taskBucket.Delete(itob(key))
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}