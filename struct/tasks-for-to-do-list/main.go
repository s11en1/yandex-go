package main

import (
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (task Task) IsOverdue() bool {
	return task.deadline.Before(time.Now())
}

func (task Task) IsTopPriority() bool {
	return task.priority >= 4
}
