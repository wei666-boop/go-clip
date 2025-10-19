package model

import "time"

type History struct {
	Id      int
	Content string
	Time    time.Time
}
