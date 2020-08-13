package controllers

// @to-do : add bcrypt to Registration
import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aabukhader/backEnd/db"
	"github.com/aabukhader/backEnd/helper"
	"github.com/aabukhader/backEnd/models"
)

var conn *sql.DB = db.ConnectDb()

func login(username string, password string) models.User {
	var user = models.User{}
	conn.QueryRow("SELECT * FROM users where username=? AND password=?", username, password).
		Scan(
			&user.ID,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Email,
		)
	return user

}

// Authenticate : check the user credentials and generat the token
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var credentials models.Authentication
	json.NewDecoder(r.Body).Decode(&credentials)
	if len(credentials.Username) == 0 || len(credentials.Password) == 0 {
		var res models.StatusRes
		res.Status = 400
		res.Msg = "username or password is invalid"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := login(credentials.Username, credentials.Password)
	if res.ID != 0 {
		token, err := helper.GetToken(credentials.Username, credentials.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var res models.StatusRes
			res.Status = 500
			res.Msg = "Error generating JWT token"
			json.NewEncoder(w).Encode(res)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Authorization", "Bearer "+token.AccessToken)
			var res models.UserStatusResSuccss
			res.Status = 200
			res.Msg = "LogedIn successfully"
			res.Data = login(credentials.Username, credentials.Password)
			res.Data.Token = token
			json.NewEncoder(w).Encode(res)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		var res models.UserStatusResFail
		res.Status = 400
		res.Msg = "username or password is invalid"
		res.Data[0] = ""
		json.NewEncoder(w).Encode(res)
	}
}

func getUserByID(id int64) models.User {
	var user = models.User{}
	conn.QueryRow("SELECT * FROM users where id=?", id).
		Scan(
			&user.ID,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Email,
		)
	defer conn.Close()
	return user
}

func createUser(user models.User) (models.User, error) {
	stmt, _ := conn.Prepare("INSERT INTO users SET username=?, password=?, firstname=?, lastname=?, email=?")
	res, err := stmt.Exec(user.Username, user.Password, user.FirstName, user.LastName, user.Email)
	id, _ := res.LastInsertId()
	defer stmt.Close()
	return getUserByID(id), err
}

// Registration : create new user
func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	if len(user.Username) == 0 || len(user.Password) == 0 || len(user.FirstName) == 0 || len(user.LastName) == 0 || len(user.Email) == 0 {
		var res models.StatusRes
		res.Status = 400
		res.Msg = "All Fields are Required"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	stmt, err := createUser(user)
	if err == nil {
		token, err := helper.GetToken(user.Username, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var res models.StatusRes
			res.Status = 500
			res.Msg = "Error generating JWT token"

			json.NewEncoder(w).Encode(res)
		} else {
			w.WriteHeader(http.StatusOK)
			var res models.UserStatusResSuccss
			res.Status = 200
			res.Msg = "User has been created successfully"
			res.Data = stmt
			res.Data.Token = token
			json.NewEncoder(w).Encode(res)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		var res models.StatusRes
		res.Status = 500
		res.Msg = "Something wont wrong"
		json.NewEncoder(w).Encode(res)
	}
}
