package controllers

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

// Authenticate : check the user credentials
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
		w.WriteHeader(http.StatusOK)
		token, err := helper.GetToken(credentials.Username)
		if err != nil {
			var res models.StatusRes
			res.Status = 500
			res.Msg = "Error generating JWT token"
			json.NewEncoder(w).Encode(res)
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
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
