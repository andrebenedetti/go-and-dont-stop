package main

import (
	"time"
)

type Task func()

func Schedule(task Task, interval time.Duration) {
	for {
		task()
		time.Sleep(interval)
	}
}