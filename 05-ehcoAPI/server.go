package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ข้อควรรู้ตัวเเปรในstruct ถ้าเป็นตัวเล็กจะเป็น private

var users = []User{
	{ID: 1, Name: "Anuchito", Age: 22},
	{ID: 2, Name: "Jimmy", Age: 28},
	{ID: 3, Name: "weerakul23May", Age: 32},
}

type Logger struct {
	Hanler http.Handler
}

type Err struct {
	Message string `json:"message"`
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUserHandler(c echo.Context) error {

	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	users = append(users, u)

	return c.JSON(http.StatusCreated, u)
}

func AuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username != "apidesign" || password != "45678" {
		fmt.Println("suererererer:", username, password)
		return false, nil
	}
	return true, nil
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //กันserver ตาย

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	g := e.Group("/users") //นึกตรงนี้ดีๆนะจ๊ะพล
	g.Use(middleware.BasicAuth(AuthMiddleware))
	g.GET("", getUsersHandler) //นึกเเล้วมาดู argument ตัวเเรก นะจ๊ะ
	g.POST("", createUserHandler)

	log.Fatal(e.Start(":2565"))
}
