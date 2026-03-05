package notes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const getHandlerName = "GET api/v1/notes?pid=?user_id="

type Note struct {
	Id   int    `json:"id"`
	Text string `json:"note"`
	Done string `json:"done"`
}

type GetNotesResponse struct {
	Notes []Note `json:"notes"`
}

func Get(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var pidStr string
	pidStr = request.URL.Query().Get("pid")
	var pid *int
	if pidStr != "" {
		value, err := strconv.Atoi(pidStr)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{Error: invalid pid}`))
			return
		}
		pid = &value
	}

	user_id := request.URL.Query().Get("user_id")

	dbConn := db.Get()
	rows, err := dbConn.Query(ctx,
		`SELECT id, note, done 
		FROM notes 
		WHERE user_id = $1 
		AND (
			(parent_id = $2) OR ($2 IS NULL AND parent_id IS NULL)
        )`,
		user_id, pid,
	)
	defer rows.Close()

	if err != nil {
		go logger.Log(err.Error(), getHandlerName)

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	var responseBody GetNotesResponse
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Id, &note.Text, &note.Done)

		if err != nil {
			go logger.Log(err.Error(), getHandlerName)

			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{Error: db error}`))
			return
		}
		responseBody.Notes = append(responseBody.Notes, note)
	}

	go logger.Log("Success", getHandlerName)
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(&responseBody)
	if err != nil {
		go logger.Log(err.Error(), getHandlerName)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: cant parse json}`))
		return
	}
}
