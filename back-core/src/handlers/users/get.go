package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const kHandlerName = "GET api/v1/users/{user_id}"

type GetUsersResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func Get(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var responseBody GetUsersResponse
	vars := mux.Vars(request)
	userId := vars["user_id"]

	if userId == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: userid cant be empty}`))
		return
	}

	dbConn := db.Get()

	err := dbConn.QueryRow(ctx,
		`SELECT username, email, name 
		FROM users
		WHERE id = $1`, userId,
	).Scan(&responseBody.Username, &responseBody.Email, &responseBody.Name)

	if err != nil {
		go logger.Log(err.Error(), kHandlerName)

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	err = json.NewEncoder(response).Encode(&responseBody)
	if err != nil {
		go logger.Log(err.Error(), kHandlerName)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}

	go logger.Log("Success", kHandlerName)
	response.WriteHeader(http.StatusOK)
}
