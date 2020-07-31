package entity

import "time"

type Session struct {
	Id        int64
	Uuid      string
	Email     string
	UserId    int64
	CreatedAt time.Time
}
