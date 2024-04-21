package utils

import (
	"encoding/json"
	"errors"
)

type StringOrSlice struct {
	Values []string
}

func (s *StringOrSlice) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	slice, ok := tmp.([]interface{})
	if ok {
		return handleSlice(slice, s)
	}
	singleString, ok := tmp.(string)
	if ok {
		s.Values = []string{singleString}
		return nil
	}
	return errors.New("field neither slice of string or string")
}

func handleSlice(slice []interface{}, s *StringOrSlice) error {
	values := []string{}
	for _, item := range slice {
		if _, ok := item.(string); !ok {
			return errors.New("field is not a slice of strings")
		}
		values = append(values, item.(string))
	}
	s.Values = values
	return nil
}
