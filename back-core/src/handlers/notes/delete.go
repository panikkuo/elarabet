package notes

import (
	"net/http"

	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/logger"
)

const deleteHandlerName = "DELETE api/v1/notes"

func Delete(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	user_id := request.URL.Query().Get("user_id")
	note_id := request.URL.Query().Get("note_id")

	if user_id == "" || note_id == "" {
		go logger.Log("Error: user id or note id empty", deleteHandlerName)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{Error: user id or note id empty}`))
		return
	}

	dbConn := db.Get()

	_, err := dbConn.Exec(ctx,
		`DELETE FROM notes
		WHERE (id = $1) AND (user_id = $2)`,
		note_id, user_id,
	)

	if err != nil {
		go logger.Log(err.Error(), deleteHandlerName)

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Error: db error}`))
		return
	}

	go logger.Log("Success", deleteHandlerName)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{Status: deleted}`))
}
