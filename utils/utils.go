package utils

import "encoding/json"

func Struct2Attr(v interface{}) (*map[string][]string, error) {
	attr := make(map[string][]string)

	_t, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(_t, &attr)
	if err != nil {
		return nil, err
	}

	return &attr, nil
}
