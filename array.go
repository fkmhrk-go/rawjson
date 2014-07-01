package rawjson

import (
	"encoding/json"
	"errors"
	"strings"
)

type RawJsonArray ([]interface{})

func (j RawJsonArray) Int(i int) (int, error) {
	v := j[i]
	if val, ok := v.(float64); ok {
		return int(val), nil
	}
	return 0, errors.New("At " + string(i) + " is not a number")
}

func (j RawJsonArray) Float(i int) (float64, error) {
	v := j[i]
	if val, ok := v.(float64); ok {
		return val, nil
	}
	return 0, errors.New("At " + string(i) + " is not a number")
}

func (j RawJsonArray) String(i int) (string, error) {
	v := j[i]
	if val, ok := v.(string); ok {
		return val, nil
	}
	return "", errors.New("At " + string(i) + " is not a string")
}

func (j RawJsonArray) Bool(i int) (bool, error) {
	v := j[i]
	if val, ok := v.(bool); ok {
		return val, nil
	}
	return false, errors.New("At " + string(i) + " is not a boolean")
}

func (j RawJsonArray) Object(i int) (RawJsonObject, error) {
	v := j[i]
	if s, ok := v.(map[string]interface{}); ok {
		return s, nil
	}
	return nil, errors.New("At " + string(i) + " is not a JsonObject")
}

func (j RawJsonArray) Array(i int) (RawJsonArray, error) {
	v := j[i]
	if s, ok := v.([]interface{}); ok {
		return s, nil
	}
	return nil, errors.New("At " + string(i) + " is not a JsonArray")
}

func ArrayFromString(v string) (RawJsonArray, error) {
	r := strings.NewReader(v)
	doc := json.NewDecoder(r)
	var obj []interface{}
	err := doc.Decode(&obj)
	if err != nil {
		return nil, err
	} else {
		return obj, nil
	}
}
