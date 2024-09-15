package main

import "fmt"

type LogLevel string

const Error LogLevel = "Error"
const Info LogLevel = "Info"

type Logger interface {
	Log()
}

type Log struct {
	Level LogLevel
}

func (logger Log) Log(errorText string) {
	switch logger.Level {
	case Error:
		fmt.Println("ERROR:", errorText)
	case Info:
		fmt.Println("INFO:", errorText)
	}
}
