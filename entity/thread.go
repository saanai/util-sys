package entity

import "time"

type Thread struct {
	Id        int64
	Uuid      string
	Topic     string
	UserId    int64
	CreatedAt time.Time
}
