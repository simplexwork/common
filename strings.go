package common

import (
	"regexp"
	"strconv"
	"strings"
)

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

// ToLower 字符串转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper 字符串转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Trim 去首位空格
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// IsEmail 判断邮箱
func IsEmail(s string) bool {
	reg := regexp.MustCompile("^[A-Za-z0-9_-\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	return reg.MatchString(s)
}

// IsMobile 判断中国手机号
func IsMobile(s string) bool {
	reg := regexp.MustCompile("^1\\d{2}\\d{8}$")
	return reg.MatchString(s)
}
