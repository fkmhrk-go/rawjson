package rawjson

import (
	"testing"
)

func TestArray_1000_int(t *testing.T) {
	val := "[123, 456]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Int(0)
	if err != nil {
		t.Errorf("Failed to Int : %s", err)
	}
	if v != 123 {
		t.Errorf("Value is %d : not 123", v)
	}
}

func TestArray_1200_float(t *testing.T) {
	val := "[123.75, 456.78]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Float(0)
	if err != nil {
		t.Errorf("Failed to Float : %s", err)
	}
	if v != 123.75 {
		t.Errorf("Value is %f : not 123.75", v)
	}
}

func TestArray_1300_string(t *testing.T) {
	val := "[\"fkm\",\"moke\"]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.String(1)
	if err != nil {
		t.Errorf("Failed to String : %s", err)
	}
	if v != "moke" {
		t.Errorf("Value is %s : not moke", v)
	}
}

func TestArray_1400_bool(t *testing.T) {
	val := "[\"fkm\", true]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Bool(1)
	if err != nil {
		t.Errorf("Failed to String : %s", err)
	}
	if !v {
		t.Errorf("Value is %d : not moke", v)
	}
}

func TestArray_1600_object(t *testing.T) {
	val := "[\"fkm\",{\"name\":\"moke\"}]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Object(1)
	if err != nil {
		t.Errorf("Failed to String : %s", err)
	}
	v2, _ := v.String("name")
	if v2 != "moke" {
		t.Errorf("Value is %s : not moke", v)
	}
}

func TestArray_1700_array(t *testing.T) {
	val := "[\"fkm\",[\"moke\"]]"
	r, err := ArrayFromString(val)
	if err != nil {
		t.Errorf("Failed to create Json : %s", err)
	}
	v, err := r.Array(1)
	if err != nil {
		t.Errorf("Failed to String : %s", err)
	}
	v2, _ := v.String(0)
	if v2 != "moke" {
		t.Errorf("Value is %s : not moke", v)
	}
}
