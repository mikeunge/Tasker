package entities

import "time"

type ITask struct {
	Id          int
	Title       string
	Description string
	Status      string
	Done        uint8
	CreatedAt   time.Time
}
