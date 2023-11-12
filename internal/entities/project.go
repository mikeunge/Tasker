package entities

import "time"

type IProject struct {
	Id   int
	Name string
}

type IDBProject struct {
	Id int
	IProject
	CreatedAt time.Time
}
