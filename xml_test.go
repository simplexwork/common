package common

import "testing"

type Student1 struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"age"`
}

func TestXML(t *testing.T) {
	str := "<Student1><name>你</name></Student1>"
	var student Student1
	err := FromXML(str, &student)
	if err != nil {
		t.Error(err)
	}
	t.Log(student.Name)

	student.Name = "你"
	str, err = ToXML(student)
	if err != nil {
		t.Error(err)
	}
	t.Log(str)
}
