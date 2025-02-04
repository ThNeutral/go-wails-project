package main

import "fmt"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (a *Logger) Log(payload string) {
	fmt.Println(payload)
}
