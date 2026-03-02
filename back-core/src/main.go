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
	"github.com/panikkuo/elarabet/back-core/src/handlers/signup"
	"github.com/panikkuo/elarabet/back-core/src/handlers/users"
	"github.com/panikkuo/elarabet/back-core/src/logger"
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
	r.HandleFunc("/v1/signup", signup.Post).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", login.Post).Methods(http.MethodPost)
	r.HandleFunc("/v1/users/{username}", users.Get).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is start!")
	_ = srv.ListenAndServe()
}
