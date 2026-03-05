package login

import (
	"encoding/json"
	"net/http"

	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const kHandlerName = "POST api/v1/login"

type PostLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostLoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func Post(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody PostLoginRequest
	var responseBody PostLoginResponse

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		go logger.Log(err.Error(), kHandlerName)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"Error": "cant parse json}"`))
		return
	}
	defer request.Body.Close()

	dbConn := db.Get()

	var rpassword string

	err = dbConn.QueryRow(ctx,
		`SELECT id, username, password 
		FROM users
		WHERE username = $1`, requestBody.Username,
	).Scan(&responseBody.Id, &responseBody.Username, &rpassword)

	if err != nil {
		go logger.Log(err.Error(), kHandlerName)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	//TODO: сделать адекватное сравнение паролей
	if rpassword != requestBody.Password {
		go logger.Log("Incorrect username or password", kHandlerName)
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{Error: username or password incorrect}`))
		return
	}

	go logger.Log("Success", kHandlerName)
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(&responseBody)
	if err != nil {
		go logger.Log(err.Error(), kHandlerName)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}
}
