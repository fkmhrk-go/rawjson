package rawjson

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

type RawJsonObject (map[string]interface{})

func (j RawJsonObject) Int(k string) (int, error) {
	v := j[k]
	if i, ok := v.(float64); ok {
		return int(i), nil
	}
	return 0, errors.New(k + " is not a number")
}

func (j RawJsonObject) Long(k string) (int64, error) {
	v := j[k]
	if i, ok := v.(float64); ok {
		return int64(i), nil
	}
	return 0, errors.New(k + " is not a number")
}

func (j RawJsonObject) Float(k string) (float64, error) {
	v := j[k]
	if i, ok := v.(float64); ok {
		return i, nil
	}
	return 0, errors.New(k + " is not a number")
}

func (j RawJsonObject) String(k string) (string, error) {
	v := j[k]
	if s, ok := v.(string); ok {
		return s, nil
	}
	return "", errors.New(k + " is not a string")
}

func (j RawJsonObject) Bool(k string) (bool, error) {
	v := j[k]
	if val, ok := v.(bool); ok {
		return val, nil
	}
	return false, errors.New(k + " is not a boolean")
}

func (j RawJsonObject) Object(k string) (RawJsonObject, error) {
	v := j[k]
	if s, ok := v.(map[string]interface{}); ok {
		return s, nil
	}
	return nil, errors.New(k + " is not a JsonObject")
}

func (j RawJsonObject) Array(k string) (RawJsonArray, error) {
	v := j[k]
	if s, ok := v.([]interface{}); ok {
		return s, nil
	}
	return nil, errors.New(k + " is not a JsonArray")
}

func ObjectFromString(v string) (RawJsonObject, error) {
	return ObjectFromReader(strings.NewReader(v))
}

func ObjectFromReader(r io.Reader) (RawJsonObject, error) {
	doc := json.NewDecoder(r)
	var obj map[string]interface{}
	err := doc.Decode(&obj)
	if err != nil {
		return nil, err
	} else {
		return obj, nil
	}
}
