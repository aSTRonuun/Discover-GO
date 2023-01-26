package server

import (
	"basic_crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser is a function that creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user

	if erro := json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("Error unmarshalling request body"))
		return
	}

	db, err := db.Connection()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into users(name, email) values(?, ?)")
	if erro != nil {
		w.Write([]byte("Error preparing statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	idInserted, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Error getting last inserted id"))
		return
	}

	// Status Code 201: Created
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with success! ID: %d", idInserted)))
}

// GetAllUsers is a function that returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connection()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	rows, erro := db.Query("select * from users")
	if erro != nil {
		w.Write([]byte("Error getting users"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if erro := rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Error scanning users"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Error encoding users to JSON"))
		return
	}
}

// GetUserById is a function that returns a user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Error converting id to integer"))
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	row, erro := db.Query("select * from users where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Error getting user"))
		return
	}

	var user user
	if row.Next() {
		if erro := row.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Error encoding user to JSON"))
		return
	}
}

// UpdateUser is a function that updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Error converting id to integer"))
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user
	if erro := json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("Error unmarshalling request body"))
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update users set name = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Error preparing statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, ID); erro != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser is a function that deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Error converting id to integer"))
	}

	db, erro := db.Connection()
	if erro != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from users where id = ?")
	if erro != nil {
		w.Write([]byte("Error preparing statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
