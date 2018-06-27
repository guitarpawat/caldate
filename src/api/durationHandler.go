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
	From string
	To string
	Result {
		Date string
		Year string
		Minute string
		Hour string
		Week string
		Percent string
	}
}

func toJSON(startDate Date, endDate Date){
	days = caldate.ResultDay(startDate, endDate)
	weeks = caldate.UnitWeek(days)
	second = caldate.convertToSecond(days)
	mins = caldate.convertToMin(second)

	response := Response({
		From: startDate
	})

}

