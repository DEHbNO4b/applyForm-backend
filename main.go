package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DEHbNO4b/applyForm/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"
	"gopkg.in/yaml.v2"
)

// var dsn string = "postgres://tykym:tykym@localhost:5432/tykym?"
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func (c *Config) parse() {
	var configPath = flag.String("config", "./config/config.yml", "path to config file")
	flag.Parse()
	configYml, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("reading config.yml error %v", err)
	}
	err = yaml.Unmarshal(configYml, c)
	if err != nil {
		log.Fatalf("can't parse congig.yml: %v", err)
	}
}

var mainConfig = Config{}

func main() {
	//считывание файла конфигурации
	mainConfig.parse()
	//соединение с БД

	dsn := "postgres://" + mainConfig.Username + ":" + mainConfig.Password + "@" + mainConfig.Host + ":" + mainConfig.Port + "/" + mainConfig.Dbname + "?"
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

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/apllye/{id}", ah.DelApply)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	s.ListenAndServe()

}
