package users

import (
	db "backend/db"
	"net/http"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	"fmt"
)



// "users
func GetAllUsers(c echo.Context) error {	
	usersDB,err := db.ReadAll()
	if err != nil {
		return c.String(http.StatusNotFound ,"Nouser found")
	}
	fmt.Print(usersDB)
	var userlist []User
	fmt.Println(userlist)
	return c.JSON(http.StatusOK ,usersDB)
}

// "user/:id"
func GetUser(c echo.Context) error {	
	id := c.Param("id")

	userDB ,err := db.Read(id)
	if err != nil {
		return c.String(http.StatusNotFound ,"not found")
	}
	user := User{
	    Name : userDB.Name,
		Email : userDB.Email,
	}
	return  c.JSON(http.StatusOK, user)
}

func Save(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	userDB := db.UserDB{
		Name : user.Name,
		Email : user.Email,
	}

	if err:= db.Create(userDB); err != nil {
		return c.String(http.StatusExpectationFailed,"create user failed")
	}

	return c.JSON(http.StatusCreated, user)
}


func Update(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "update ="+id)
}


func Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "delete ="+id)
}