package model

import (
	"time"
)

type Link struct {
	Id            uint64    `json:"id" gorm:"primaryKey"`
	OriginLink    string    `json:"origin_link"`
	ShortenedLink string    `json:"shortened_link"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateLink struct {
	OriginLink    string `json:"origin_link"`
	ShortenedLink string `json:"shortened_link"`
}

type UpdateLink struct {
	OriginLink string `json:"origin_link"`
}
