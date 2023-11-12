package entities

import "time"

// Custom types
type TEvent string
type TLogLevel string
type TEntity string

const (
	// Event (Action)
	ELogEventModified TEvent = "modified"
	ELogEventCreated  TEvent = "created"
	ELogEventDeleted  TEvent = "deleted"
	// Log level
	ELogLevelInfo    TLogLevel = "info"
	ELogLevelWarning TLogLevel = "warn"
	ELogLevelError   TLogLevel = "error"
	// Entity - Which entity created the log
	ELogEntityTask    TEntity = "task"
	ELogEntityProject TEntity = "project"
)

type ILog struct {
	Id        int
	Event     TEvent
	Level     TLogLevel
	Entity    TEntity
	Caller    string
	Message   string
	CreatedAt time.Time
}
