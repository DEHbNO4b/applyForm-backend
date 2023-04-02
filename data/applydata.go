package data

import (
	"database/sql"
	"encoding/json"
	"io"
)

type ApplyData struct {
	Id              string `json:"apply_id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	FathersName     string `json:"fathers_name,omitempty"`
	BorneDate       string `json:"borne_date,omitempty"`
	Adress1         string `json:"adress1,omitempty"`
	PassportSeries  string `json:"passport_series,omitempty"`
	PassportNumber  string `json:"passport_number,omitempty"`
	DateIssue       string `json:"date_issue,omitempty"`
	PropertyType    string `json:"property_type,omitempty"`
	PropertyNumber1 string `json:"property_number1,omitempty"`
	PropertyNumber2 string `json:"property_number2,omitempty"`
	Adress2         string `json:"adress2,omitempty"`
	Amount          string `json:"amount,omitempty"`
	Date            string `json:"date,omitempty"`
}

func (ad *ApplyData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ad)
}
func (ad *ApplyData) AddApplyData(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO applydata (firstName, lastName, fathersName, bornedate, adress1, passportSeries, passportNumber,
									dateIssue,propertyType,propertyNumber1,propertyNumber2,adress2,amount,date)
						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14);
					`, ad.FirstName, ad.LastName, ad.FathersName, ad.BorneDate, ad.Adress1, ad.PassportSeries, ad.PassportNumber,
		ad.DateIssue, ad.PropertyType, ad.PropertyNumber1, ad.PropertyNumber2, ad.Adress2, ad.Amount, ad.Date)
	if err != nil {
		return err
	}
	return nil
}
