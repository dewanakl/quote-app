package model

import "gorm.io/gorm"

type Quotes struct {
	Anime string `json:"anime"`
	Char  string `json:"character"`
	Quote string `json:"quote"`
	gorm.Model
}

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}
