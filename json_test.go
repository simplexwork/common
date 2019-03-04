package common

import "testing"

type Student struct {
	Name string `json:"name"`
}

func TestJSON(t *testing.T) {
	str := "{\"name\":\"我\"}"
	var student Student
	err := FromJSON(str, &student)
	if err != nil {
		t.Error(err)
	}
	t.Log(student.Name)

	student.Name = "你"
	str, err = ToJSON(student)
	if err != nil {
		t.Error(err)
	}
	t.Log(str)
}
