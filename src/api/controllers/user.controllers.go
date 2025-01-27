package controllers

import (
	"encoding/json"
	"go-crud/src/api/models"
	"go-crud/src/config"
	"go-crud/src/database"
	"net/http"
	"strconv"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.ConnectToPostgres()

	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}

	users, err := database.GetUsersQuery(db)
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" || newUser.Age == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectToPostgres()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = database.InsertUserQuery(db, newUser.Name, newUser.Age)
	if err != nil {
		http.Error(w, "Unable to insert user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.ConnectToPostgres()

	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}

	var userId string = r.URL.Query().Get("id")
	if userId == "" {
		http.Error(w, "Error 404", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserQuery(db, userId)
	if err != nil {
		http.Error(w, "Unable to fetch the user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.ConnectToPostgres()

	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}

	var userId string = r.URL.Query().Get("id")
	if userId == "" {
		http.Error(w, "Error 404", http.StatusBadRequest)
		return
	}

	var userData models.User
	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	userIdInt, err := strconv.Atoi(userId)
	userData.Id = userIdInt
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := database.PutUserQuery(db, userId, userData.Name, userData.Age)
	if err != nil {
		http.Error(w, "Unable to update user.", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.ConnectToPostgres()

	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}

	var userId string = r.URL.Query().Get("id")
	if userId == "" {
		http.Error(w, "Error 404", http.StatusBadRequest)
		return
	}

	user, err := database.DeleteUserQuery(db, userId)
	if err != nil {
		http.Error(w, "Unable to fetch the user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
