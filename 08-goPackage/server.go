package main

import (
	"database/sql"
	"log"

	"github.com/JimmyTanapon/go-package/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Connection string
	db = user.InitDB()
	defer db.Close()
	log.Print("ok!")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //กันserver ตาย

	e.GET("/users", user.GetUsersHandler(db))
	e.GET("/users/:id", user.GetUsertByIdHandler(db))
	e.POST("/users", user.CreateUserHandler(db))

	log.Fatal(e.Start(":2565"))

	// createTb := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );`

	// _, err = db.Exec(createTb)
	// if err != nil {
	// 	log.Fatal("can't create table ", err)

	// }
	// fmt.Println("Create sucess!")
}

// อยากได้ของ ใช้ db.QueryRow
// ไม่อยากได้ของ ใช้db.Exec
