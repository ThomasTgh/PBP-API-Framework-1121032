package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/praktikum/Modul_Explorasi/controllers"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", controllers.GetUsers) //Read
	e.POST("/users", controllers.InsertUsers) //Create
	e.PUT("/users/", controllers.UpdateUsers) //Update
	e.DELETE("/users/:id", controllers.DeleteUsers) //Delete

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //fitur CORS
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"}, //list link yang diijinkan
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}, //list method yang diijinkan
	}))

	e.Logger.Fatal(e.Start(":8080"))
}