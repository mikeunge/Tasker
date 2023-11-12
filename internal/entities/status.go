package entities

import "time"

type IStatus struct {
	Id        int
	Title     string
	CreatedAt time.Time
}
