package todo

import (
	"encoding/json"
	"log"
	"time"
)

type Todo struct {
	TodoId   uint64     `json:"id"`
	Name     string     `json:"name"`
	DoneDate *time.Time `json:"donedate"`
}

func NewTodo(id uint64, name string) Todo {
	return Todo{
		TodoId:   id,
		Name:     name,
		DoneDate: nil,
	}
}

func (t *Todo) Do() {
	now := time.Now()
	t.DoneDate = &now
}

func (t Todo) IsDone() bool {
	return t.DoneDate != nil
}

func (t Todo) WasDoneInDate(d time.Time) bool {
	if t.DoneDate == nil {
		return false
	}
	y1, m1, d1 := t.DoneDate.Date()
	y2, m2, d2 := d.Date()
	return (y1 == y2) && (m1 == m2) && (d1 == d2)
}

func (t Todo) AsJsonByteArr() []byte {
	ret, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
