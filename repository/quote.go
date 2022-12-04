package repository

import "gorm.io/gorm"

type Quote struct {
	db *gorm.DB
}

func NewQuoteRepository(g *gorm.DB) *Quote {
	return &Quote{db: g}
}
