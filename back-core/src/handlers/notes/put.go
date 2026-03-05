package notes

import (
	"encoding/json"
	"net/http"

	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const updateHandlerName = "PUT api/v1/notes"

type UpdateRequestBody struct {
	UserId string  `json:"user_id"`
	NoteId int     `json:"note_id"`
	Note   *string `json:"note"`
	Done   *int    `json:"done"`
}

func Put(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody UpdateRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		go logger.Log(err.Error(), updateHandlerName)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}
	defer request.Body.Close()

	if requestBody.Note != nil {
		dbConn := db.Get()
		_, err = dbConn.Exec(ctx,
			`UPDATE notes
			SET note = $1
			WHERE id = $2`,
			requestBody.Note, requestBody.NoteId,
		)

		if err != nil {
			go logger.Log(err.Error(), updateHandlerName)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{Error: db error}`))
			return
		}
	}

	if (requestBody.Done != nil) && (*requestBody.Done != 0 && *requestBody.Done != 1) {
		go logger.Log("Error: done must be 0, 1", updateHandlerName)

		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: done must be 0, 1"}`))
		return
	}

	if requestBody.Done != nil {
		dbConn := db.Get()
		_, err = dbConn.Exec(ctx,
			`UPDATE notes
			SET done = $1
			WHERE id = $2`,
			requestBody.Done, requestBody.NoteId,
		)

		if err != nil {
			go logger.Log(err.Error(), updateHandlerName)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{Error: db error}`))
			return
		}
	}

	go logger.Log("Success", updateHandlerName)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{Status: edited}`))
}
