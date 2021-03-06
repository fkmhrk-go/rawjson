package rawjson

import (
	"testing"
)

func TestObject_0000_int(t *testing.T) {
	val := "{\"name\":\"fkm\",\"score\":100}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Int("score")
	if err != nil {
		t.Errorf("Failed to GetInt : %s", err)
	}
	if v != 100 {
		t.Errorf("Value is not %d", v)
	}
}

func TestObject_0100_long(t *testing.T) {
	val := "{\"name\":\"fkm\",\"time\":1422902874000}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Long("time")
	if err != nil {
		t.Errorf("Failed to Get Long : %s", err)
	}
	if v != 1422902874000 {
		t.Errorf("Value is not %d", v)
	}
}

func TestObject_0100_float(t *testing.T) {
	val := "{\"name\":\"fkm\",\"rate\":4.56}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Float("rate")
	if err != nil {
		t.Errorf("Failed to GetFloat : %s", err)
	}
	if v != 4.56 {
		t.Errorf("Value is not %f", v)
	}
}

func TestObject_0200_string(t *testing.T) {
	val := "{\"name\":\"fkm\",\"rate\":4.56}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.String("name")
	if err != nil {
		t.Errorf("Failed to GetString : %s", err)
	}
	if v != "fkm" {
		t.Errorf("Value is not %s", v)
	}
}

func TestObject_0300_bool(t *testing.T) {
	val := "{\"name\":\"fkm\",\"enabled\":true}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Bool("enabled")
	if err != nil {
		t.Errorf("Failed to Bool : %s", err)
	}
	if !v {
		t.Errorf("Value is %d : not true", v)
	}
}

func TestObject_0300_Object(t *testing.T) {
	val := "{\"name\":\"fkm\",\"ext\":{\"data\":\"123\"}}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Object("ext")
	if err != nil {
		t.Errorf("Failed to GetJsonObject : %s", err)
	}
	v2, _ := v.String("data")
	if v2 != "123" {
		t.Errorf("Value is not %s", v2)
	}
}

func TestObject_0500_Array(t *testing.T) {
	val := "{\"name\":\"fkm\",\"ext\":[\"123\"]}"
	r, err := ObjectFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v2, err := r.Array("ext")
	if err != nil {
		t.Errorf("Failed to JsonArray : %s", err)
	}
	v3, err := v2.String(0)
	if err != nil {
		t.Errorf("Failed to String of Array : %s", err)
	}
	if v3 != "123" {
		t.Errorf("Value is %s, not 123", v3)
	}
}
