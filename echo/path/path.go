package path

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}


func Show(c echo.Context) error {
	// Get team and member from the query string
	// http://localhost:1323/show?team=x-men&member=wolverine => test url
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}


func Save(c echo.Context) error {
	// Get name and email
	// curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save => command for run this path
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}



//command for post data struct
// curl -X POST http://localhost:1323/users \
//      -H "Content-Type: application/json" \
//      -d '{"name": "John Doe", "email": "john@example.com"}'

func Users(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
}