package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a main page")
	tem, err := template.ParseFiles("resources/index.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func agents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using an agents page")
	tem, err := template.ParseFiles("resources/agents.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a login page")
	tem, err := template.ParseFiles("resources/login.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func support(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a support page")
	tem, err := template.ParseFiles("resources/support.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func main() {
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./resources/static"))))
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/agents", agents)
	http.HandleFunc("/support", support)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
