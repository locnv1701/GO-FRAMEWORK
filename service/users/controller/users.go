package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"base/pkg/db"
	"base/pkg/log"
	"base/pkg/router"
	"base/service/users/model"
)

// resGetUsers Struct
type resGetUsers struct {
	Status  bool         `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
}

// GetUser Function to Get All User Data
func GetUser(w http.ResponseWriter, r *http.Request) {

	// Database Query here
	rows, err := db.PSQL.Query("SELECT * FROM users")
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}
	var user model.User
	var users []model.User

	// Populate Data
	for rows.Next() {
		// Match / Binding Database Field with Struct
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			// Print Log Error
			log.Println("error", "get-user", err.Error())
		} else {
			// Append User Struct to Users Array of Struct
			users = append(users, user)
		}
	}
	defer rows.Close()

	var response resGetUsers

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Write Response Data to HTTP
	router.ResponseWrite(w, response.Code, response)
}

// AddUser Function to Add User Data
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Database Query
	query := fmt.Sprintf("INSERT INTO users VALUES "+"(%d, '%s', '%s');", user.ID, user.Name, user.Email)
	log.Println(log.LogLevelDebug, "user-query", query)

	_, err := db.PSQL.Exec(query)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseCreated(w)
}

// GetUserByID Function to Get User Data By User ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Database Query
	query := fmt.Sprintf("SELECT * FROM users WHERE id=%d", userID)
	rows, err := db.PSQL.Query(query)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	var user model.User
	var users []model.User

	// Populate Data
	for rows.Next() {
		// Match / Binding Database Field with Struct
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			// Print Log Error
			log.Println("error", "get-user", err.Error())
		} else {
			// Append User Struct to Users Array of Struct
			users = append(users, user)
		}
	}
	defer rows.Close()

	var response resGetUsers

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Write Response Data to HTTP
	router.ResponseWrite(w, response.Code, response)
}

// PutUserByID Function to Update User Data By User ID
func PutUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	var user model.User

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Database Query
	query := fmt.Sprintf("UPDATE users SET username='%s', email='%s' WHERE id=%d", user.Name, user.Email, userID)
	log.Println(log.LogLevelDebug, "user-query", query)

	_, err = db.PSQL.Exec(query)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseUpdated(w)
}

// DelUserByID Function to Delete User Data By User ID
func DelUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Database Query
	query := fmt.Sprintf("DELETE FROM users WHERE id=%d ", userID)
	log.Println(log.LogLevelDebug, "user-query", query)

	_, err = db.PSQL.Query(query)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "")
}
