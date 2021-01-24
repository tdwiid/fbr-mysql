package book

import (
    "fmt"

    "github.com/elliotforbes/go-fiber-tutorial/database"
    "github.com/gofiber/fiber"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
    gorm.Model
    Title  string `json:"name"`
    Author string `json:"author"`
    Rating int    `json:"rating"`
}