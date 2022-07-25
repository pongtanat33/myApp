package users

import (
	db "goapi/db"
	"net/http"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	"fmt"
)



// "users
func GetAllUsers(c echo.Context) error {	
	// id := c.Param("id")
	// db.ReadAll()
	usersDB,err := db.ReadAll()
	if err != nil {
		return c.String(http.StatusNotFound ,"Nouser found")
	}
	fmt.Print(usersDB)
	// myuser := *usersDB
	// var interfaceSlice []interface{} = make([]interface{}, len(usersDB))
	// for i, d := range usersDB{
	// 	interfaceSlice[i] = d
	// }

	var userlist []User
	// userlist2 := User(usersDB)
	 
	// for _, a := range *usersDB {
	// 	// fmt.Printf("Type: %T Value: %v\n", a, a.Email)
	// 	userlist = append(userlist, User{
	// 		Email:    a.Email,
	// 		Name:     a.Name,
	// 	})
    // }

	fmt.Println(userlist)
	// result, _ := json.Marshal(usersDB)
	// fmt.Println(string(result))
	// myuser := []User{}
	// myuser = make(usersDB)
	
	// myuser := []User{
	// 	Name : usersDB.Name,
	// 	Email : usersDB.Email,
	// }
	// return  c.JSON(http.StatusOK, )
	// return c.JSON(http.StatusOK ,userlist)
	return c.JSON(http.StatusOK ,usersDB)
	// return c.String(http.StatusOK,"readall")	
}

// "user/:id"
func GetUser(c echo.Context) error {	
	id := c.Param("id")

	userDB ,err := db.Read(id)
	if err != nil {
		return c.String(http.StatusNotFound ,"not found")
	}
	user := User{
		// Id : userDB.Code,
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
		// Code : user.Id,
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