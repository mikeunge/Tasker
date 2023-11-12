package entities

import "time"

// Custom types
type TAction string
type TCaller string

const (
	// Action types
	ELogActionModified TAction = "modified"
	ELogActionCreated  TAction = "created"
	ELogActionDeleted  TAction = "deleted"
	// Entity - Which entity created the log
	ELogCallerProject TCaller = "project"
	ELogCallerTask    TCaller = "task"
	ELogCallerNotes   TCaller = "notes"
	ELogCallerStatus  TCaller = "status"
)

type ILog struct {
	Id        int
	Code      string
	Action    TAction
	Caller    TCaller
	Message   string
	CreatedAt time.Time
}
