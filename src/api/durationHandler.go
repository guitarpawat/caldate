package api

import (
	"caldate"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	minute := (caldate.ConvertToMin(second))
	hour := caldate.ConvertToHour(minute)
	week := caldate.UnitWeek(day)
	percent := caldate.CalPercent(day)
	percentStr := ""

	dayComma := toComma(uint64(day))
	secondComma := toComma(second)
	minuteComma := toComma(minute)
	hourComma := toComma(hour)

	if startDate.Year == endDate.Year {
		percentStr = fmt.Sprintf("%.2f%% of %d", percent, startDate.Year)
	} else {
		percentStr = fmt.Sprintf("%.2f%% of a common year (365 days)", percent)
	}
	jsonStr, err := toJSON(from, to, fmt.Sprintf("%s days", dayComma), "years",
		fmt.Sprintf("%s seconds", secondComma), fmt.Sprintf("%s minutes", minuteComma),
		fmt.Sprintf("%s hours", hourComma), week, percentStr)
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

func toComma(n uint64) string{
    in := []byte(strconv.FormatUint(n, 10))
    var out []byte
    if i := len(in) % 3; i != 0 {
        if out, in = append(out, in[:i]...), in[i:]; len(in) > 0 {
            out = append(out, ',')
        }
    }
    for len(in) > 0 {
        if out, in = append(out, in[:3]...), in[3:]; len(in) > 0 {
            out = append(out, ',')
        }
    }
    return string(out)
}