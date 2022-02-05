package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	UID       `gorm:"primaryKey"`
	Title     string
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewTask(title string, text string) (*Task, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return &Task{}, err
	}

	if title == "" || text == "" {
		err = fmt.Errorf("title and text cannot be empty")
		return &Task{}, err
	}

	var task = &Task{
		Title:     title,
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	task.SetId(uid)
	return task, nil
}

func (task *Task) GetId() string {
	return task.Id.String()
}
