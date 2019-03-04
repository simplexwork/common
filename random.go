package common

import (
	"math/rand"
	"time"
)

// RandomString 随机字符串
func RandomString(le int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < le; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomNumber 随机数字
func RandomNumber(le int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < le; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
