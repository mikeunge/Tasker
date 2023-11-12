package entities

import "time"

type IProject struct {
	Id        int
	Name      string
	CreatedAt time.Time
}
