package past

import "time"

func OneHourAgo() time.Time {
	return time.Now().Add(-time.Hour)
}
