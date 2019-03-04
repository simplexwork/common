package common

import (
	"testing"
)

func TestAes(t *testing.T) {
	key := []byte("1234567890123456")
	os := []byte(`测试测试`)
	b, err := AesEncrypt(os, key)
	if err != nil {
		t.Error(err)
		return
	}
	str := Base64Encode(b)
	t.Logf("crypto: %v", str)

	bb, err := Base64Decode(str)
	if err != nil {
		t.Error(err)
		return
	}
	b1, err := AesDecrypt(bb, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("decrypt: %v", string(b1))

}
