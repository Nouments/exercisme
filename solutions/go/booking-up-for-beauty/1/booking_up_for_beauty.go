package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a time.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return hour >= 12 && hour < 24 && t.Second()>0 && t.Minute()>=0
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, at %s.",
		t.Format("Monday, January 2, 2006"),
		t.Format("15:04"))
}


func AnniversaryDate() time.Time {
	now := time.Now()
	anniversary := time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.Local)
	return anniversary
}
