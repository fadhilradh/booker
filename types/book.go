package types

type NewBook struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Price int `json:"price" binding:"required,number"`
}

type Publisher struct {
	Name string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}