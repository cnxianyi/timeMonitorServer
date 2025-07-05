package _type

import "time"

type Time struct {
	Title   string    `json:"title" binding:"required"`
	Process string    `json:"process" binding:"required"`
	Time    time.Time `json:"time" binding:"required"`
}

type Times struct {
	Data []Time `json:"data" binding:"required"`
}
