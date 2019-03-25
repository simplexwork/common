package common

import "encoding/json"

// ToJSON 对象转json字符串
func ToJSON(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON json字符串转对象
func FromJSON(data []byte, i interface{}) error {
	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}
	return nil
}
