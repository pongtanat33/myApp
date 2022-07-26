package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	u "backend/users"

	t "backend/todos"

	db "backend/db"
)

func main() {
	db.DB()
	db.Migrate()
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,`{"response" : "success"}`)
	})

	
	e.GET("users",u.GetAllUsers)
	e.GET("users/:id",u.GetUser)
	e.POST("users",u.Save)
	e.PUT("users/:id",u.Update)
	e.DELETE("users/d/:id",u.Delete)

	e.GET("todos/:email",t.GetTodo)
	e.POST("todos",t.Save)
	e.PUT("todos",t.Update)
	e.DELETE("todos/d/:id",t.Delete)
	
	e.Logger.Fatal(e.Start(":1234"))
}



