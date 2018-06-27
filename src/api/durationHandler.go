package api

import (
	"caldate"
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
	endtDate := caldate.NewDate(r.FormValue("EndDate"),
		r.FormValue("EndMonth"),
		r.FormValue("EndYear"))
	fmt.Fprintf(w, "%d", caldate.ResultDay(startDate, endtDate))
}

func doAction(startDate, endDate caldate.Date, r *http.Request) {
	from := caldate.FormatDateConverter(startDate)
	to := caldate.FormatDateConverter(endDate)
	day := caldate.ResultDay(startDate, endDate)
	second := caldate.ConvertToSecond(day)
	minute := caldare.ConvertToMin(second)
	week := caldate.UnitWeek(day)
	response := toJSON(from, to, day, "", second, minute, "", week, "")
}

type Response struct {
	From string `json:"from"`
	To string `json:"to"`
	Date string `json:"date"`
	Years string `json:"years"`
	Seconds string `json:"seconds"`
	Minutes string `json:"minutes"`
	Hours string `json:"hours"`
	Weeks string `json:"weeks"`
	Percent string `json:"percent"`
}

func toJSON(from, to, dates, years, seconds, minutes, hours, weeks, percent string) string{
	response := Response({
		From: from,
		To: to,
		Dates: dates
		Years: years
		Seconds: seconds
		Minutes: minutes
		Hours:  hours
		Weeks: weeks
		Percent: percent
	})
	json, err := json.Marshal(response)
    return (string(json))
}
