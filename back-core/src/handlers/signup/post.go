package signup

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const kHandlerName = "POST api/v1/signup"

type PostSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func Post(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody PostSignupRequest

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		go logger.Log(err.Error(), kHandlerName)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}
	defer request.Body.Close()

	dbConn := db.Get()

	_, err = dbConn.Exec(ctx,
		`INSERT INTO users(username, email, password, name) 
		VALUES($1, $2, $3, $4)`,
		requestBody.Username, requestBody.Email, requestBody.Password, requestBody.Name,
	)

	if err != nil {
		go logger.Log(err.Error(), kHandlerName)

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			response.WriteHeader(http.StatusConflict)
			response.Write([]byte(`{Error: Username or email already exists}`))
			return
		}

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	go logger.Log("Success", kHandlerName)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{Status: User was added}`))
}
