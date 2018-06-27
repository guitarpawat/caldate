package api

import (
	"caldate"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), 500)
	}
	startDate := caldate.NewDate(r.FormValue("StartDate"),
		r.FormValue("StartMonth"),
		r.FormValue("StartYear"))
	endDate := caldate.NewDate(r.FormValue("EndDate"),
		r.FormValue("EndMonth"),
		r.FormValue("EndYear"))
	doAction(startDate, endDate, w)
}

func doAction(startDate, endDate caldate.Date, w http.ResponseWriter) {
	from := caldate.FormatDateConverter(startDate)
	to := caldate.FormatDateConverter(endDate)
	day := caldate.ResultDay(startDate, endDate)
	second := caldate.ConvertToSecond(day)
	minute := caldate.ConvertToMin(second)
	week := caldate.UnitWeek(day)
	jsonStr, err := toJSON(from, to, fmt.Sprintf("%d", day), "years",
		fmt.Sprintf("%d", second), fmt.Sprintf("%d", minute), "hours",
		week, "percent")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), 500)
		return
	}
	fmt.Fprint(w, jsonStr)
}

type Response struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Dates   string `json:"date"`
	Years   string `json:"years"`
	Seconds string `json:"seconds"`
	Minutes string `json:"minutes"`
	Hours   string `json:"hours"`
	Weeks   string `json:"weeks"`
	Percent string `json:"percent"`
}

func toJSON(from, to, dates, years, seconds, minutes, hours, weeks, percent string) (string, error) {
	response := Response{
		From:    from,
		To:      to,
		Dates:   dates,
		Years:   years,
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
		Weeks:   weeks,
		Percent: percent,
	}
	json, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return "", err
	}
	return (string(json)), nil
}
