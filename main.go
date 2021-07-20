package main

import (
	"calender/dateControl"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", dataHandler)
	http.ListenAndServe(":8080", nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	days := dateControl.GetThisMonthDays(t)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := tmpl.Execute(w, days); err != nil {
		panic(err.Error())
	}
}
