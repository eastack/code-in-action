package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id       int
	Uuid     string
	Body     string
	UserId   int
	ThreadId int
	CreateAt time.Time
}
