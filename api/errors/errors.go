package errors

import "errors"

var (
	TaskNotFound       = errors.New("the task does not exist")
	NoTasksFound       = errors.New("there are no tasks in the db")
	FailedToAddTask    = errors.New("could not create new task")
	FailedToUpdateTask = errors.New("failed to update task")
	EmptyTitle         = errors.New("title cannot be empty")
	EmptyText          = errors.New("text cannot be empty")
)
