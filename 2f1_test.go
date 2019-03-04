package common

import "testing"

func Test2FA(t *testing.T) {
	var ti int64 = 3
	t.Logf("time:%v crypto:%v", ti, TFA([]byte("test"), ti))
}
