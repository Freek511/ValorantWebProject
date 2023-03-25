package controllers

import (
	"ValorantWebProject/internal/models"
	"ValorantWebProject/internal/services"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Handlers interface {
	Home(w http.ResponseWriter, r *http.Request)
	Agents(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	CheckUser(w http.ResponseWriter, r *http.Request)
	Support(w http.ResponseWriter, r *http.Request)
}

type PageHandlers struct {
	userService services.Service
}

func (p *PageHandlers) Home(w http.ResponseWriter, r *http.Request) {
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

func (p *PageHandlers) Agents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using an agents page")
	tem, err := template.ParseFiles("resources/agents.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func (p *PageHandlers) CheckUser(w http.ResponseWriter, r *http.Request) {
	//p.Login(w, r)
	var body []byte
	r.Body.Read(body)
	var user models.User
	json.Unmarshal(body, &user)
	status, err := p.userService.CheckUser(user.Name, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
	}
	if status {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
}

func (p *PageHandlers) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a login page")
	tem, err := template.ParseFiles("resources/login.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func (p *PageHandlers) Support(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using a support page")
	tem, err := template.ParseFiles("resources/support.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}
