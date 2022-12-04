package db

import "gorm.io/gorm"

type Quotes struct {
	Anime string `json:"anime"`
	Char  string `json:"character"`
	Quote string `json:"quote"`
	gorm.Model
}
