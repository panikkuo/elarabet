package signup

import (
	"encoding/json"
	"net/http"

	"github.com/panikkuo/elarabet/back-core/src/db"
)

const kAddUser = ""

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type SignupResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func V1Signup(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody SignupRequest
	var responseBody SignupResponse

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(response, "Error: cant parse json", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	dbConn := db.Get()

	err = dbConn.QueryRow(ctx,
		"SELECT id, username FROM add_user($1, $2, $3, $4)",
		requestBody.Username, requestBody.Email, requestBody.Password, requestBody.Name,
	).Scan(&responseBody.Id, &responseBody.Username)

	if err != nil {
		http.Error(response, "Error: db error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
