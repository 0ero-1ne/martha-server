package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	Id           uint            `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Email        string          `json:"email"`
	Password     string          `json:"-"`
	Username     string          `json:"username"`
	Image        string          `json:"image"`
	Role         string          `json:"role"`
	SavedBooks   SavedBooks      `json:"saved_books"`
	Comments     []Comment       `json:"comments,omitempty"`
	BookRates    []BooksRates    `json:"book_rates,omitempty"`
	CommentRates []CommentsRates `json:"comment_rates,omitempty"`
}

type SavedBook struct {
	BookId    uint   `json:"book_id"`
	ChapterId uint   `json:"chapter_id"`
	Page      int    `json:"page"`
	Audio     uint64 `json:"audio"`
}

type SavedBooks map[BookFolder][]SavedBook

func (savedBooks *SavedBooks) Value() (driver.Value, error) {
	return json.Marshal(savedBooks)
}

func (savedBooks *SavedBooks) Scan(value any) error {
	content, err := value.([]byte)
	if !err {
		return errors.New("failed to read json data in SavedBooks")
	}

	return json.Unmarshal(content, &savedBooks)
}

type BookFolder string

const (
	Reading   BookFolder = "Reading"
	Ended     BookFolder = "Ended"
	Stopped   BookFolder = "Stopped"
	Planed    BookFolder = "Planed"
	Favorites BookFolder = "Favorites"
)
