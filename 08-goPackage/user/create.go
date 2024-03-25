package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		// Handler logic using db
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

}
