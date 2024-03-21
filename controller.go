package controller

import (

	//"fmt"
	"log"
	m "modul4/Model"
	"net/http"

	"github.com/labstack/echo/v4"
	//"github.com/gorilla/mux"
)

// GET AllUsers
func GetAllUsers(c echo.Context) error {
	log.Print("test log print")
	db := connect()
	defer db.Close()

	query := "SELECT * FROM Users "

	name := c.QueryParam("name")
	age := c.QueryParam("age")

	if name != "" {
		query += "WHERE name='" + name + "'"
	}

	if age != "" {
		if name != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age=" + age
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return sendErrorResponse(c)
	}
	defer rows.Close()

	var users []m.User
	for rows.Next() {
		var user m.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			log.Println(err)
			return sendErrorResponse(c)
		}
		users = append(users, user)
	}

	return sendSuccessResponseAll(c, users)
}

// Insert
func InsertUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	var newUser m.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := db.Exec("INSERT INTO users(ID, Name, Age, Address) VALUES (?, ?, ?, ?)",
		newUser.ID,
		newUser.Name,
		newUser.Age,
		newUser.Address,
	)

	if err != nil {
		log.Println(err)
		return sendErrorResponse(c)
	}

	return sendSuccessResponse(c, newUser)
}

// delete
func DeleteUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	userID := c.Param("user_id")

	_, err := db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		log.Println(err)
		return sendErrorResponse(c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}

// sendSuccessResponse sends a success response
func sendSuccessResponse(c echo.Context, newUser m.User) error {
	response := m.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    newUser,
	}
	return c.JSON(http.StatusOK, response)
}

func sendSuccessResponseAll(c echo.Context, user []m.User) error {
	response := m.UsersResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}

// sendErrorResponse sends an error response
func sendErrorResponse(c echo.Context) error {
	response := m.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "Failed",
	}
	return c.JSON(http.StatusBadRequest, response)
}
