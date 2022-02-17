package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/mikeunge/Tasker/api/errors"
)

// Database model
type Task struct {
	Id        string `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Done      bool   `json:"done"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Functionality for this entity
type TaskRepository interface {
	GetAll() []Task
	Get(string) (*Task, error)
	Add(string, string) (string, error)
	Update(string, string, string) error
	Delete(*Task) error
}

func NewTask(title, text string, done bool) (*Task, error) {
	if title == "" {
		return &Task{}, errors.EmptyTitle
	}
	if text == "" {
		return &Task{}, errors.EmptyText
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return &Task{}, err
	}

	task := &Task{
		Id:        uid.String(),
		Title:     title,
		Text:      text,
		Done:      done,
		CreatedAt: time.Now(),
	}
	return task, nil
}
