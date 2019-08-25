package model

import (
	"time"
)

// Archieve Object
type Archieve struct {
	ID        string     `json:"id"`
	Link      string     `json:"link"`
	Code      string     `json:"code"`
	Published string     `json:"published"`
	Journals  []*Journal `json:"journals"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
}
