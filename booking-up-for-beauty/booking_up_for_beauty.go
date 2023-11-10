package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	formattedTime, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}
	return formattedTime

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	formattedTime, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return formattedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	formattedTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return formattedTime.Hour() >= 12 && formattedTime.Hour() < 18

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	formattedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("You have an appointment on %s", formattedTime.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
