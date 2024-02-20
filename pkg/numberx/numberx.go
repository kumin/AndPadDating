package numberx

import "strconv"

func ParseInt(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}
