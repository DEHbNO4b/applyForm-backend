package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DEHbNO4b/applyForm/data"
	"github.com/gorilla/mux"
)

type Apply struct {
	l  *log.Logger
	db *sql.DB
}

// var tpl = template.Must(template.ParseFiles("./web/build/index.html"))

func NewApply(l *log.Logger, db *sql.DB) *Apply {
	return &Apply{l: l, db: db}
}

func (f Apply) GetApply(rw http.ResponseWriter, r *http.Request) {
	//f.l.Println("in get apply method")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
	}

	row := f.db.QueryRow(`SELECT firstName, lastName, fathersName, borneDate, adress1,
								 passportSeries, passportNumber, dateIssue, propertyType,
								 propertyNumber1, propertyNumber2, amount, adress2, date from applydata where id =$1`, id)

	if row.Err() != nil {
		f.l.Println(row.Err())
		http.Error(rw, "unable to get apply data", http.StatusInternalServerError)
	}
	a := data.ApplyData{}
	row.Scan(&a.FirstName, &a.LastName, &a.FathersName, &a.BorneDate, &a.Adress1,
		&a.PassportSeries, &a.PassportNumber, &a.DateIssue, &a.PropertyType,
		&a.PropertyNumber1, &a.PropertyNumber2, &a.Amount, &a.Adress2, &a.Date)
	f.l.Println(a)
	data, err := json.Marshal(a)
	if err != nil {
		f.l.Println(err)
		http.Error(rw, "unable to murshal json", http.StatusInternalServerError)
	}
	rw.Write(data)
}
func (f Apply) GetApplyes(rw http.ResponseWriter, r *http.Request) {
	var applyCollection []data.ApplyData
	rows, err := f.db.Query(`SELECT id, firstName, lastName, fathersName, date from applydata`)
	if err != nil {
		f.l.Println(err)
		//http.Error(rw, "query db err", http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		a := data.ApplyData{}
		if err = rows.Scan(&a.Id, &a.FirstName, &a.LastName, &a.FathersName, &a.Date); err != nil {
			f.l.Println(err)
			http.Error(rw, "query db err", http.StatusInternalServerError)
		}
		applyCollection = append(applyCollection, a)
	}
	data, err := json.Marshal(applyCollection)
	if err != nil {
		f.l.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(data)

}
func (f Apply) DelApply(rw http.ResponseWriter, r *http.Request) {
	f.l.Println("in delete apply method")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
	}

	_, err = f.db.Exec(`DELETE FROM applydata where id =$1`, id)

	if err != nil {
		f.l.Println(err)
		http.Error(rw, "unable to delete apply data", http.StatusInternalServerError)
	}

}

func (f Apply) PostApplyes(rw http.ResponseWriter, r *http.Request) {
	apply := data.ApplyData{}
	err := apply.FromJSON(r.Body)
	if err != nil {
		f.l.Println(err)
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	err = apply.AddApplyData(f.db)
	if err != nil {
		f.l.Println(err)
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

}
