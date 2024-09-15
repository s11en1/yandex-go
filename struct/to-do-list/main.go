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

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (list ToDoList) TasksCount() int {
	return len(list.tasks)
}

func (list ToDoList) NotesCount() int {
	return len(list.notes)
}

func (list ToDoList) CountTopPrioritiesTasks() int {
	count := 0

	for _, task := range list.tasks {
		if task.IsTopPriority() {
			count++
		}
	}

	return count
}

func (list ToDoList) CountOverdueTasks() int {
	count := 0

	for _, task := range list.tasks {
		if task.IsOverdue() {
			count++
		}
	}

	return count
}
