package api

import (
	"caldate"
	"fmt"
	"net/http"
	"strconv"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), 500)
	}
	startDate := caldate.Date{Date: toi(r.FormValue("StartDate")),
		Month: toi(r.FormValue("StartMonth")),
		Year:  toi(r.FormValue("StartYear"))}
	endtDate := caldate.Date{Date: toi(r.FormValue("EndDate")),
		Month: toi(r.FormValue("EndMonth")),
		Year:  toi(r.FormValue("EndYear"))}
	fmt.Fprintln(w, "%d", caldate.ResultDay(startDate, endtDate))
}

func toi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

Type Response struct {
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
		From: from
		To: to
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
