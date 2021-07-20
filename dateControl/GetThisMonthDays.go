package dateControl

import (
	"calender/domain"
	"time"
)

func GetThisMonthDays(t time.Time) domain.OneMonthDays {
	days := domain.OneMonthDays{
		Year:  int(t.Year()),
		Month: int(t.Month()),
		Days:  [6][7]int{},
	}

	first := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	last := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
	previous_last := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)

	day_counter := 1

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if j < int(first.Weekday()) && i == 0 {
				days.Days[i][j] = previous_last.Day() - (int(first.Weekday()) - j - 1)
			} else {
				if day_counter == last.Day()+1 {
					day_counter = 1
				}
				days.Days[i][j] = day_counter
				day_counter++
			}
		}
	}

	return days
}
