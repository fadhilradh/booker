package types

import "time"

type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
	Rating int `json:"rating"`
	Year int `json:"year"`
	Lang string `json:"lang"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}