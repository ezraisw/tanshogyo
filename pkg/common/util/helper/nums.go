package helper

import "strconv"

func AssumeInt(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		// Parse as float to allow for decimals.
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return int(f)
		}
		return 0
	default:
		return 0
	}
}
