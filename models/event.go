package models

import "time"

type Event struct {
	ID          int
	Name        string      `required:"true"`
	Description string      `required:"true"`
	Location    string      `required:"true"`
	DateTime    time.Time   `required:"true"`
	UserID      int
}

var events = []Event{}

func Save(e Event) {

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
