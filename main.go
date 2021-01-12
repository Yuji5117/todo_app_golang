package main

import (
	"strconv"

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
	db, err := gorm.Open("sqlite3", "test.sqlite3")
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
	todos := dbGetAll()
	ctx.HTML(200, "index.html", gin.H{"todos": todos})
}

var postValue = func (ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbInsert(text, status)
	ctx.Redirect(302, "/")
}

var getDetail = func (ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := dbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
}

var updateValue = func (ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

var deleteConfirm = func (ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := dbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
}

var deleteValue = func (ctx *gin.Context) {
	n := ctx.Param("id")
	id , err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	dbDelete(id)
	ctx.Redirect(302, "/")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	// index
	router.GET("/", getValue)

	// create
	router.POST("/create", postValue)

	// detail
	router.GET("/detail/:id", getDetail)

	// update
	router.POST("/update/:id", updateValue)

	// delete confirm
	router.GET("/delete_confirm/:id", deleteConfirm)

	// delete
	router.POST("/delete/:id", deleteValue)

	router.Run()
}


