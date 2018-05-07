package logger

import "time"

func GetJpTimeNow() time.Time {
	return time.Now().In(getJpTimeLocation())
}

func getJpTimeLocation() *time.Location {
	timezone := "Asia/Tokyo"
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.FixedZone(timezone, 9*60*60)
	}

	return loc
}
