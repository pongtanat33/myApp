package todos

import (
	db "goapi/db"
	"net/http"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	"fmt"
	"strconv"
)



// "users
func GetAllTodos(c echo.Context) error {	
	// id := c.Param("id")
	// db.ReadAll()
	usersDB,err := db.ReadAll()
	if err != nil {
		return c.String(http.StatusNotFound ,"No user found")
	}
	fmt.Print(usersDB)
	// result, _ := json.Marshal(usersDB)
	// fmt.Println(string(result))
	// myuser := Todo{
	// 	Id : userDB.Code,
	// 	Name : userDB.Name,
	// 	Email : userDB.Email,
	// }
	// return  c.JSON(http.StatusOK, )
	return c.JSON(http.StatusOK ,usersDB)
	// return c.String(http.StatusOK,"readall")	
}

// "todo/:id"
func GetTodo(c echo.Context) error {	
	Email := c.Param("email")

	todoDB ,err := db.ReadTodo(Email)
	if err != nil {
		return c.String(http.StatusNotFound ,"not found")
	}

	return  c.JSON(http.StatusOK, todoDB)
}

func Save(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	todoDB := db.TodoDB{
		Complete : todo.Complete,
		Description : todo.Description,
		UserEmail : todo.UserEmail,
	}

	if err:= db.CreateTodo(todoDB); err != nil {
		return c.String(http.StatusExpectationFailed,"create todo failed")
	}

	return c.JSON(http.StatusCreated, todo)
}


func Update(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	todoDB := db.TodoDB{
		ID:todo.ID,
		Description : todo.Description,
		UserEmail : todo.UserEmail,
		Complete : todo.Complete,
	}
	
	if err:= db.UpdateTodo(todoDB); err != nil {
		return c.String(http.StatusExpectationFailed,"update todo failed")
	}

	return c.JSON(http.StatusCreated, todoDB)
}


func Delete(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	id :=c.Param("id")
	fmt.Println("id",id)
	uid, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        fmt.Println("err",err)
    }
	// fmt.Println(id)
	todoDB := db.TodoDB{
		ID:uid,
		Description : todo.Description,
		UserEmail : todo.UserEmail,
		Complete : todo.Complete,
	}
	
	if err:= db.DeleteTodo(todoDB); err != nil {
		return c.String(http.StatusExpectationFailed,"delete todo failed")
	}

	return c.JSON(http.StatusCreated, todo)
}