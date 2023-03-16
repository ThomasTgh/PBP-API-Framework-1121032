package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := connect()
	defer db.Close()

	rows,err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User
	var users []User
	for rows.Next(){
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
			log.Fatal(err)
		}else{
			users = append(users, user)
		}
	}
	return c.JSON(http.StatusOK, users)
}

func InsertUsers(c echo.Context) error {
	db := connect()
	defer db.Close()

	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	UserType, _ := strconv.Atoi(c.FormValue("type"))
	
	stmt, err := db.Prepare("INSERT INTO users(name, age, address, user_type) values (?,?,?,?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer stmt.Close()

	_, errQuery := stmt.Exec(name, age, address, UserType)
	
	var response UserResponse
	if errQuery == nil{
		response.Status = 200
		response.Message = "Success"
	}else{
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Insert Failed!"
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUsers(c echo.Context) error{
	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	UserType, _ := strconv.Atoi(c.FormValue("type"))

	stmt, err := db.Prepare("UPDATE users SET name='?', age='?', address='?', user_type='?' WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer stmt.Close()

	_, errQuery := stmt.Exec(name, age, address, UserType, id)

	var response UserResponse
	if errQuery == nil{
		response.Status = 200
		response.Message = "Success"
	}else{
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Update Failed!"
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUsers(c echo.Context) error{
	db := connect()
	defer db.Close()

	id := c.Param("id")

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?", id)

	var response UserResponse
	if errQuery == nil{
		response.Status = 200
		response.Message = "Success"
	}else{
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed!"
	}

	return c.JSON(http.StatusOK, response)
}