package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DEHbNO4b/applyForm/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"
)

var dsn string = "postgres://postgres:917836@localhost:5432/lightning?"

func main() {
	//соединение с БД
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS applydata(
		id serial primary key,
		firstName varchar(20),
		lastName varchar(20),
		fathersName varchar(20),
		borneDate varchar(20),
		adress1 text,
		passportSeries varchar(20),
		passportNumber varchar(20),
		dateIssue varchar(20),
		propertyType varchar(20),
		propertyNumber1 varchar(20),
		propertyNumber2 varchar(20),
		amount varchar(20),
		adress2 text,
		date varchar(20)

	) `)
	if err != nil {
		panic(err)
	}
	l := log.New(os.Stdout, "products-api: ", log.LstdFlags)
	ah := handlers.NewApply(l, db)
	sm := mux.NewRouter()

	// startPage := sm.Methods(http.MethodGet).Subrouter()
	// startPage.HandleFunc("/", ah.GetStartPage)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/applyes", ah.GetApplyes)
	getRouter.HandleFunc("/applye/{id}", ah.GetApply)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/applyes", ah.PostApplyes)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	s.ListenAndServe()

}
