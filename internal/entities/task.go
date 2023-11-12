package entities

import "time"

type ITask struct {
	Id          int
	ProjectId   int
	Title       string
	Description string
	StatusId    int
	CreatedAt   time.Time
}
