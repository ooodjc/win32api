package win32api

import "strconv"

func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
