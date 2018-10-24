package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

// Task is the struct representing the todo task
type Task struct {
	Key   int
	Value string
}

func (t *Task) String() string {
	return fmt.Sprintf("%d. %s", t.Key, t.Value)
}

// Store represents the BoltDB storage
type Store struct {
	db     *bolt.DB
	bucket []byte
}

// Get retrieves the task at the specified key
func (s *Store) Get(key int) (*Task, error) {
	var ret *Task
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		v := b.Get(itob(key))
		ret.Key = key
		ret.Value = string(v)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// GetAll retrieves all of the tasks
func (s *Store) GetAll() ([]Task, error) {
	var tasks []Task
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		c := b.Cursor()
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

// Put adds the given data and returns the key for that data
func (s *Store) Put(task string) (int, error) {
	var id int
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

// Delete removes the task at the given key
func (s *Store) Delete(key int) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		return b.Delete(itob(key))
	})
}

// New creates a new BoltDB Store
func New(path, bucket string) (*Store, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	if err != nil {
		return nil, err
	}

	s := &Store{
		db:     db,
		bucket: []byte(bucket),
	}

	return s, nil
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
