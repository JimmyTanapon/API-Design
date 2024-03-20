package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "postgres"   // as defined in docker-compose.yml
	password = "jumonji123" // as defined in docker-compose.yml
	dbname   = "users"      // as defined in docker-compose.yml
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb
	defer sdb.Close()
	log.Print("ok!")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //กันserver ตาย

	e.GET("/users", getUsersHandler)
	e.GET("/users/:id", getUsertByIdHandler)
	e.POST("/users", createUserHandler)

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

func getUsersHandler(c echo.Context) error {
	users, err := queryAll(db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, users)
}

func getUsertByIdHandler(c echo.Context) error {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := queryOneRow(db, userid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)

}

func createUserHandler(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	id, err := insertUsers(db, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	response := map[string]interface{}{
		"Message": "user create success",
		"id":      id,
	}

	return c.JSON(http.StatusOK, response)
}
