package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_"github.com/mattn/go-sqlite3"
)


type Todo struct {
	gorm.Model
	Text string
	Status string
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "rest.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed!!!")
	}
	var todo Todo
	db.Order("created_at desc").Find(&todo, id)
	db.Close()
	return todo
}

var getValue = func (ctx *gin.Context) {
	data := "Hello /Go"
	ctx.HTML(200, "index.html", gin.H{"data": data})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", getValue)

	router.Run()
}

