package common

import "testing"

func TestSequence(t *testing.T) {
	idWoker := NewIDWorker(1)

	c := make(chan ID)

	cnt := 100000

	for i := 0; i < cnt; i++ {
		go func() {
			id, err := idWoker.Next()
			if err != nil {
				t.Error(err)
			}
			c <- id
		}()
	}

	m := make(map[ID]int)
	for i := 0; i < cnt; i++ {
		id := <-c
		t.Log(id)
		_, ok := m[id]
		if ok {
			t.Error("exists!!")
		}
		m[id] = i
	}
	t.Log("ok")
}
