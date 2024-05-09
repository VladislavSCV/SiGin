package models

type Auto struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Year int `json:"year"`
	Price int `json:"price"`
}
