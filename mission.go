package main

import "time"

type mission struct {
	Name            string
	CreateTime      time.Time
	DeadLineTime    time.Time
	CompleteTime    time.Time
	TaskDescription string
	Comment         string
}
