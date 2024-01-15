package past

import "time"

// careful: not 'time.Now()'
var Now = time.Now

func OneHourAgo() time.Time {
	return Now().Add(-time.Hour)
}
