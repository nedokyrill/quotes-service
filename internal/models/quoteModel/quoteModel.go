package quoteModel

import "github.com/google/uuid"

type Quote struct {
	Id     uuid.UUID `json:"id"`
	Author string    `json:"author"`
	Text   string    `json:"quote"`
}
