package day_of_the_week

import "time"

func dayOfTheWeek(day, month, year int) string {
	week := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
	return week
}
