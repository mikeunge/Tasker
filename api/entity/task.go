package entity

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

var (
	errorEmptyTitle = errors.New("title cannot be empty")
	errorEmptyText  = errors.New("text cannot be empty")
)

// Database model
type Task struct {
	base      `gorm:"primaryKey"`
	Title     string
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Functionality for this entity
type TaskRepository interface {
	GetAll() []Task
	Get(string) (*Task, error)
	Add(Task) (string, error)
	Update(Task) error
}

func NewTask(title string, text string) (*Task, error) {
	if title == "" {
		return &Task{}, errorEmptyTitle
	}
	if text == "" {
		return &Task{}, errorEmptyText
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return &Task{}, err
	}

	var t = &Task{
		Title:     title,
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	t.base.SetId(uid.String())

	return t, nil
}
