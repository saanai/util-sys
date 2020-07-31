package entity

import "time"

type Post struct {
	Id        int64
	Uuid      string
	Body      string
	UserId    int64
	ThreadId  int64
	CreatedAt time.Time
}
