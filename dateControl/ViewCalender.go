package dateControl

import (
	"fmt"
	"time"
)

func ViewCalender(t time.Time) {
	days := GetThisMonthDays(t)
	for _, week := range days.Days {
		fmt.Println(week)
	}
}
