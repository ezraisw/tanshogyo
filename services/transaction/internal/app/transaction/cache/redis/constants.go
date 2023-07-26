package cacheredis

import "time"

const (
	FormatKey = "cart:%s"

	DurationTTL = 30 * 24 * time.Hour
)
