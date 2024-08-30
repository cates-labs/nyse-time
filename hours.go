package nysetime

import "time"

// The function checks if a given time falls within normal trading hours of the stock market.
func IsNormalHours(t time.Time) bool {
	openTime := time.Date(t.Year(), t.Month(), t.Day(), 14, 30, 0, 0, time.UTC)
	earlyCloseTime := time.Date(t.Year(), t.Month(), t.Day(), 18, 00, 0, 0, time.UTC)
	closeTime := time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, time.UTC)

	if IsEarlyClose(t) {
		return t.After(openTime) && t.Before(earlyCloseTime)
	}

	return t.After(openTime) && t.Before(closeTime)
}

func IsExtendedHours(t time.Time) bool {
	loc, _ := time.LoadLocation("America/New_York")
	earlyCloseTime := time.Date(t.Year(), t.Month(), t.Day(), 13, 00, 0, 0, loc)
	closeTime := time.Date(t.Year(), t.Month(), t.Day(), 16, 00, 0, 0, loc)
	extendedCloseTime := time.Date(t.Year(), t.Month(), t.Day(), 20, 00, 0, 0, loc)

	if IsEarlyClose(t) {
		return t.After(earlyCloseTime) && t.Before(extendedCloseTime)
	}

	return t.After(closeTime) && t.Before(extendedCloseTime)

}
