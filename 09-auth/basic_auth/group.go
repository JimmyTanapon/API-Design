package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/golang-jwt/jwt/v5"
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

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "jonh" || password != "password!" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"jon Snow",
		"admin",
		"accessToken",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := token.SignedString([]byte("secret")) //access token
	if err != nil {
		return err

	}
	// Create refresh token
	refreshTokenClaims := &jwtCustomClaims{
		"Jon Snow",
		"admin",
		"refresh",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Longer-lived
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}
func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationToken := c.Request().Header.Get("Authorization")
		if authorizationToken == "" {
			fmt.Println("1")

			return echo.ErrUnauthorized
		}
		parts := strings.Split(authorizationToken, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			fmt.Println("2")

			return echo.ErrUnauthorized
		}
		//parts คือการ เเยก Bearer ออกไป  ex Bearer tokenxxxxxx.xxxxxxx
		jwtToken := parts[1]

		token, err := jwt.ParseWithClaims(jwtToken, &jwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			fmt.Println("3")

			return echo.ErrUnauthorized
		}

		claim, ok := token.Claims.(*jwtCustomClaims)
		if !ok {
			fmt.Println("4")

			return echo.ErrUnauthorized

		}
		fmt.Println("User name:", claim)
		c.Set("user", token)
		return next(c)

	}

}

// refreshToken function
func refreshToken(c echo.Context) error {
	refreshTokenString := c.FormValue("refresh_token")

	jwtSecretKey := []byte("secret")

	// Parse the token
	token, err := jwt.ParseWithClaims(refreshTokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid || claims.Type != "refresh" {
		return echo.ErrUnauthorized
	}

	// Create new access token
	newAccessTokenClaims := &jwtCustomClaims{
		claims.Name,
		claims.Role,
		"access",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessTokenClaims)
	newAccessTokenString, err := newAccessToken.SignedString(jwtSecretKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  newAccessTokenString,
		"refresh_token": refreshTokenString,
	})
}
func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	g := e.Group("/users") //นึกตรงนี้ดีๆนะจ๊ะพล
	g.POST("/login", login)
	g.POST("/refresh", refreshToken)
	// g.Use(middleware.BasicAuth(AuthMiddleware))
	// g.Use(jwtMiddleware)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	g.Use(echojwt.WithConfig(config)) //ถ้าทำ JWTของ echo

	g.GET("", getUsersHandler) //นึกเเล้วมาดู argument ตัวเเรก นะจ๊ะ
	g.POST("", createUserHandler)

	log.Fatal(e.Start(":2565"))

}
