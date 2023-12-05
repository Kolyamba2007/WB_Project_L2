package main

import "net/http"

var (
	events = make(map[int]Event, 100)
)

func main() {
	initEvents()

	http.HandleFunc("/events_for_day", events_for_day)
	http.HandleFunc("/events_for_week", events_for_week)
	http.HandleFunc("/events_for_month", events_for_month)
	http.HandleFunc("/create_event", create_event)
	http.HandleFunc("/update_event", update_event)

	http.ListenAndServe(":8090", nil)
}
