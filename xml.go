package common

import "encoding/xml"

// ToXML 对象转xml字符串
func ToXML(obj interface{}) (string, error) {
	b, err := xml.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromXML xml字符串转对象
func FromXML(str string, obj interface{}) error {
	err := xml.Unmarshal([]byte(str), obj)
	if err != nil {
		return err
	}
	return nil
}
