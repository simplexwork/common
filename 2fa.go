package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"time"
)

// TFA : 2 Factor Authentication,双因子验证
func TFA(value []byte, t int64) string {
	// 当前时间
	epochSeconds := time.Now().Unix() / t

	// 当前时间转byte[]
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		a := epochSeconds >> shift
		b := a & mask
		result = append(result, byte(b))
	}

	// 加密
	mac := hmac.New(sha1.New, result)
	mac.Write(value)
	hash := mac.Sum(nil)

	// 20个字节转6位数字
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := (uint32(hashParts[0]) << 24) + (uint32(hashParts[1]) << 16) +
		(uint32(hashParts[2]) << 8) + uint32(hashParts[3])
	pwd := number % 1000000

	// 不足6位补零
	pwdStr := fmt.Sprintf("%d", pwd)
	for len(pwdStr) < 6 {
		pwdStr = "0" + pwdStr
	}
	return pwdStr
}
