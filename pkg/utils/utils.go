package utils

import "strconv"

func StrToInt(val string) int {
	if result, err := strconv.Atoi(val); err != nil {
		return 0
	} else {
		return result
	}
}
