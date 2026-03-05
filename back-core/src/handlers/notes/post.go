package notes

import (
	"encoding/json"
	"net/http"

	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const postHandlerName = "POST api/v1/notes"

type PostNotesRequest struct {
	UserId   string `json:"user_id"`
	ParentId *int   `json:"parent_id"`
	Note     string `json:"note"`
}

func Post(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody PostNotesRequest

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		go logger.Log(err.Error(), postHandlerName)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}
	defer request.Body.Close()

	dbConn := db.Get()
	_, err = dbConn.Exec(ctx,
		`INSERT INTO notes(parent_id, user_id, note, done) 
		VALUES ($1, $2, $3, 0)`,
		requestBody.ParentId, requestBody.UserId, requestBody.Note,
	)

	if err != nil {
		go logger.Log(err.Error(), postHandlerName)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	go logger.Log("Success", postHandlerName)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{Status: User was added}`))
}
