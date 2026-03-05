package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/panikkuo/elarabet/back-core/src/db"
	"github.com/panikkuo/elarabet/back-core/src/handlers/login"
	"github.com/panikkuo/elarabet/back-core/src/handlers/notes"
	"github.com/panikkuo/elarabet/back-core/src/handlers/signup"
	"github.com/panikkuo/elarabet/back-core/src/handlers/users"
	"github.com/panikkuo/elarabet/back-core/src/logger"
	"github.com/panikkuo/elarabet/back-core/src/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db.Init(dsn)
	logger.Log(dsn, "main")

	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}).Methods(http.MethodOptions)

	r.HandleFunc("/v1/signup", signup.Post).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", login.Post).Methods(http.MethodPost)
	r.HandleFunc("/v1/users/{user_id}", users.Get).Methods(http.MethodGet)

	r.HandleFunc("/v1/notes", notes.Post).Methods(http.MethodPost)
	r.HandleFunc("/v1/notes", notes.Get).Methods(http.MethodGet)
	r.HandleFunc("/v1/notes", notes.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/v1/notes", notes.Put).Methods(http.MethodPut)

	handlerWithCors := middleware.Cors(r)
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      handlerWithCors,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is start!")
	_ = srv.ListenAndServe()
}
