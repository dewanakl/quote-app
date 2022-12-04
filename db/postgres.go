package db

import (
	"fmt"
	"quoteapp/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	g *gorm.DB
}

func NewDB(host, username, password, database string, port int) *Postgres {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, username, password, database, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Postgres{g: conn}
}

func (p *Postgres) Migrate() {
	p.g.AutoMigrate(&model.Quotes{})
}

func (p *Postgres) DB() *gorm.DB {
	return p.g
}
