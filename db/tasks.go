package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("task")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Initconn(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return (err)
	})
}

func CreateTask(task string) (int, error) {
	var id64 int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		key := itob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id64, nil
}

func itob(val int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(64))
	return b
}

func btoi(val []byte) int {
	return int(binary.BigEndian.Uint64(val))
}
