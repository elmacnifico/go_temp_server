package main

import (
	"time"
)

type Input struct {
	Host  string
	ID    int64
	Time  time.Time
	Value float64
}
