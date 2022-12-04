package model

import (
	"quoteapp/db"

	"gorm.io/gorm"
)

type Quote struct {
	gorm *gorm.DB
}

func NewQuoteModel(g *gorm.DB) *Quote {
	return &Quote{gorm: g}
}

func (q *Quote) Insert(qoute any) error {
	return q.gorm.Create(qoute).Error
}

func (q *Quote) Get() (db.Quotes, error) {
	result := db.Quotes{}
	err := q.gorm.First(&result).Error
	return result, err
}
