package common

import "encoding/json"

// ToJSON 对象转json字符串
func ToJSON(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON json字符串转对象
func FromJSON(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		return err
	}
	return nil
}
