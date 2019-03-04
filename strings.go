package common

import "strconv"

// StrToInt 字符串转整型
func StrToInt(s string, defInt int) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = defInt
	}
	return
}

// StrToInt64 字符串转整型
func StrToInt64(s string, defInt int64) (i int64) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = defInt
	}
	return
}

// IntToStr 整型转字符串
func IntToStr(i int) (s string) {
	s = strconv.Itoa(i)
	return
}

// Int64ToStr 64位整型转字符串
func Int64ToStr(i int64) (s string) {
	s = strconv.FormatInt(i, 10)
	return
}
