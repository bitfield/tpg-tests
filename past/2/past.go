package past

import "time"

func OneHourAgo(tm time.Time) time.Time {
	return tm.Add(-time.Hour)
}
