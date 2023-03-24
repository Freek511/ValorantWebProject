package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a main page")
	tem, err := template.ParseFiles("resources/index.html")
	if err != nil {
		log.Println(err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	tem.Execute(w, nil)
}

func Agents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using an agents page")
	tem, err := template.ParseFiles("resources/agents.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a login page")
	tem, err := template.ParseFiles("resources/login.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func Support(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a support page")
	tem, err := template.ParseFiles("resources/support.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}
