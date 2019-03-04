package common

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5 .
func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}
