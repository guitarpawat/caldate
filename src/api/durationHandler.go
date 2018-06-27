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
	fmt.Fprintln(w, "%d", caldate.ResultDay(startDate, endtDate))
}

func doAction(startDate, endDate caldate.Date, r *http.Request) {
	from := caldate.FormatDateConverter(startDate)
	to := caldate.FormatDateConverter(endDate)
	diff := caldate.ResultDay(startDate, endDate)
	_, _, _ = from, to, diff
}
