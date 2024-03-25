package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		// Handler logic using db
		users, err := queryAll(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, users)

	}

}

func GetUsertByIdHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Handler logic using db
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

}
