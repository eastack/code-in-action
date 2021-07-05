package util

import "strconv"

func S2i(str string) int {
	i64, _ := strconv.ParseInt(str, 10, 32)
	return int(i64)
}
