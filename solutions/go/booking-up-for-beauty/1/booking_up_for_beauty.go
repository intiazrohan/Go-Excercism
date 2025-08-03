package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    const layout = "1/02/2006 15:04:05"
    const formattedLayout = "2006-01-02 15:04:05 +0000 UTC"
    res,_ :=time.Parse(layout,date)
	return res
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
    res,_ :=time.Parse(layout,date)
	return res.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
    res,_ :=time.Parse(layout,date)
	return (res.Hour()> 11 && res.Hour() <18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    const layout = "1/2/2006 15:04:05"
    res,_ :=time.Parse(layout,date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", res.Weekday().String() , res.Month(), res.Day(), res.Year(), res.Hour(), res.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
}
