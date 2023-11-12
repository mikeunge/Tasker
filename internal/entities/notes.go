package entities

import "time"

type INotes struct {
	Id          int
	TaskId      int
	Title       string
	Description string
	CreatedAt   time.Time
}
