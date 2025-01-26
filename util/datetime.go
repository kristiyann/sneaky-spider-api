package util

import "time"

func DatePtrEqual(date1, date2 *time.Time) bool {
	if date1 == nil || date2 == nil {
		return false
	}

	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}
