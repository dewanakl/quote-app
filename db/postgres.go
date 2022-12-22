package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	g *gorm.DB
}

func NewDB() *Postgres {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Postgres{g: conn}
}

func (p *Postgres) DB() *gorm.DB {
	p.migrate()
	return p.g
}

func (p *Postgres) migrate() {
	p.g.AutoMigrate(&Quotes{})
}
