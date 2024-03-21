package main

import (
	controller "modul4/Controller"
	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// router := mux.NewRouter()
	// router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	// router.HandleFunc("/users-insert", controller.InsertUser).Methods("POST")
	// //router.HandleFunc("/users/insert/gorm", controller.InsertUserGorm).Methods("POST")
	// // router.HandleFunc("/users", controller.DeleteUser)
	// router.HandleFunc("/users/{user_id}", controller.DeleteUser).Methods("DELETE")

	// // router.HandleFunc("/transaction", controller.GetAllTransaction).Method("GET")
	// // router.HandleFunc("/transaction", controller.InsertTransaction).Method("POST")
	// // router.HandleFunc("/transaction", controller.DeleteTransaction).Method("PUT")
	// // router.HandleFunc("/transaction/{transaction_id}", controller.DeleteTransaction).Method("PUT")

	// http.Handle("/", router)
	// fmt.Println("Connected to port 8888")
	// log.Println("Connected to port 8888")
	// log.Fatal(http.ListenAndServe(":8888", router))

	//ECHO

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define API routes
	e.GET("/users", controller.GetAllUsers)
	e.POST("/users-insert", controller.InsertUser)
	e.DELETE("/users/:user_id", controller.DeleteUser)

	// Start the HTTP server
	port := ":8888"
	e.Logger.Fatal(e.Start(port))

}
