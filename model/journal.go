package model

import (
	"time"
)

// Journal Object
type Journal struct {
	ID         string    `json:"id"`
	ArchieveID string    `json:"archieveId"`
	Title      string    `json:"title"`
	Authors    string    `json:"authors"`
	Abstract   string    `json:"abstract"`
	Link       string    `json:"link"`
	PDFLink    string    `json:"pdfLink"`
	Published  string    `json:"published"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
