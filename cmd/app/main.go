package main

import (
	"ValorantWebProject/internal/controllers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./resources/static"))))
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/agents", controllers.Agents)
	http.HandleFunc("/support", controllers.Support)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
