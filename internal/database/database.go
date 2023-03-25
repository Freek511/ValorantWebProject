package database

import (
	"ValorantWebProject/internal/models"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Repository interface {
	Config(dbName string)
	Connecting()
	GetAllUsers() ([]models.User, error)
	UserExists(username, password string) (bool, error)
	UserByID(id int) (models.User, error)
	AddUser(user models.User) (int64, error)
}

type DataBase struct {
	db  *sql.DB
	cfg mysql.Config
}

func (database *DataBase) Config(dbName string) {
	database.cfg = mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}
}

func (database *DataBase) Connecting() {
	var err error
	database.db, err = sql.Open("mysql", database.cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := database.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func (database *DataBase) GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := database.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role); err != nil {
			return nil, fmt.Errorf("GetAllUsers: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (database *DataBase) UserExists(username, password string) (bool, error) {
	var user models.User
	row := database.db.QueryRow("SELECT * FROM users WHERE name = ? and password = ?", username, password)
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("Login %d: no such user", username)
		}
		return false, fmt.Errorf("Login %d: %v", username, err)
	}
	return true, nil
}

func (database *DataBase) UserByID(id int) (models.User, error) {
	var user models.User
	row := database.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("UserByID %d: no such user", id)
		}
		return user, fmt.Errorf("UserByID %d: %v", id, err)
	}
	return user, nil
}

func (database *DataBase) AddUser(user models.User) (int64, error) {
	result, err := database.db.Exec("INSERT INTO users (email, name, password, role) VALUES (?, ?, ?, ?)",
		user.Email, user.Name, user.Password, user.Role)

	if err != nil {
		return 1, fmt.Errorf("AddUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 1, fmt.Errorf("AddUser: %v", err)
	}
	return id, nil
}
