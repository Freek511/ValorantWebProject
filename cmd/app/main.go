package main

import (
	"ValorantWebProject/internal/controllers"
	"ValorantWebProject/internal/database"
	"fmt"
	"net/http"
)

func main() {
	var db database.DataBase

	db.Config("usersdb")
	db.Connecting()
	users, err := db.GetAllUsers()
	fmt.Println("All users: ", users, "\nFound error: ", err)
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./resources/static"))))
	var handle controllers.PageHandlers
	http.HandleFunc("/", handle.Home)
	http.HandleFunc("/login", handle.Login)
	http.HandleFunc("/check_user", handle.CheckUser)
	http.HandleFunc("/agents", handle.Agents)
	http.HandleFunc("/support", handle.Support)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
