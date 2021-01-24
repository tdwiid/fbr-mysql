package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tdwiid/fbr-mysql/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

type ConnInfo struct {
	MyUser   string
	Password string
	Host     string
	Port     int
	Db       string
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}

func initDatabase() {
	var err error
	database.DBConn, err = DbConn(cn.MyUser, cn.Password, cn.Host, cn.Db, cn.Port)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	cn := ConnInfo{
		"root",
		"newpass",
		"127.0.0.1",
		3306,
		"xd_data",
	}

	db := DbConn(cn.MyUser, cn.Password, cn.Host, cn.Db, cn.Port)

	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
