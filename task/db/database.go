package db

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"task/todo"
	"time"

	"github.com/boltdb/bolt"
)

var once sync.Once
var database *bolt.DB = nil

func Open() {
	once.Do(func() {
		openedb, err := bolt.Open("TaskData.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		database = openedb
	})
}

func Close() {
	database.Close()
}

func CreateTodo(name string) error {
	return database.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		task := todo.NewTodo(id, name).AsJsonByteArr()

		return bucket.Put(itob(id), task)
	})
}

func GetTodos() ([]todo.Todo, error) {
	var tasks []todo.Todo
	err := database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tasks"))
		if bucket == nil {
			return errors.New("Bucket tasks not found")
		}
		cursor := bucket.Cursor()

		for _, bytes := cursor.First(); bytes != nil; _, bytes = cursor.Next() {
			var task todo.Todo
			json.Unmarshal(bytes, &task)
			if !task.IsDone() {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	return tasks, err
}

func RemoveTodo(id uint64) error {
	return database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tasks"))
		if bucket == nil {
			return errors.New("Could not find tasks bucket")
		}
		return bucket.Delete(itob(id))
	})
}

func CompleteTodo(id uint64) error {
	return database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tasks"))
		if bucket == nil {
			return errors.New("Could not find tasks bucket")
		}
		id := itob(id)
		taskBytes := bucket.Get(id)
		if taskBytes == nil {
			return errors.New("Could not find the todo with the requested id")
		}
		var task todo.Todo
		json.Unmarshal(taskBytes, &task)
		task.Do()
		return bucket.Put(id, task.AsJsonByteArr())
	})
}

func GetCompletedToday() ([]todo.Todo, error) {
	var tasks []todo.Todo
	err := database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tasks"))
		if bucket == nil {
			return errors.New("Could not find tasks bucket")
		}
		c := bucket.Cursor()
		for _, taskBytes := c.First(); taskBytes != nil; _, taskBytes = c.Next() {
			var task todo.Todo
			json.Unmarshal(taskBytes, &task)
			if task.WasDoneInDate(time.Now()) {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	return tasks, err
}

func itob(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}
