package times

import (
	"fmt"
	"time"
)

func Times() {
	now := time.Now()                                   // Get time now
	fmt.Println(now.Year(), now.Month(), now.Day())     // Year, Month, Day
	fmt.Println(now.Hour(), now.Minute(), now.Second()) // Hour, Minute, Second
}
