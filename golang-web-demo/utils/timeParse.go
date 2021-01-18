package utils

import "time"

func ParseDateTime(thetime string) (result time.Time) {
	const dateTimeForm = "2006-01-02 15:04:05"
	result, _ = time.Parse(dateTimeForm, thetime)
	return
}

func ParseDate(thetime string) (result time.Time) {
	const dateForm = "2006-01-02"
	result, _ = time.Parse(dateForm, thetime)
	return
}

func MaxDateInList(list []string) (maxDate string) {
	tmp := ParseDate("1900-01-01")
	for _, thedate := range list {
		result := ParseDate(thedate)
		if result.After(tmp) {
			tmp = result
		}
	}
	maxDate = tmp.Format("2006-01-02")
	return
}
