package jsonutil

import (
	"strings"
	"testing"
)

type JsonObjTest struct {
	stringField string `json:"string_field"`
	boolField   bool   `json"bool_field"`
}

func TestReadObjectFromJSONReader(t *testing.T) {
	testConfigJsonContent := `{
	"string_field": "string_value",
	"bool_field": true
}`

	var jsonObjTest JsonObjTest
	if err := ReadObjectFromJSONReader(strings.NewReader(testConfigJsonContent), &jsonObjTest); err != nil {
		t.Error("Failed to read: ", err)
	}
	if jsonObjTest.stringField != "string_value" {
		t.Errorf("Read invalid - stringField doesn't match. Expected: string_value | Got: %s", jsonObjTest.stringField)
	}
	if jsonObjTest.boolField != true {
		t.Errorf("Read invalid - boolField doesn't match. Expected: true | Got: %s", jsonObjTest.boolField)
	}
}

func TestReadObjectFromJSONString(t *testing.T) {
	testConfigJsonContent := `{
	"string_field": "string_value",
	"bool_field": true
}`

	var jsonObjTest JsonObjTest
	if err := ReadObjectFromJSONString(testConfigJsonContent, &jsonObjTest); err != nil {
		t.Error("Failed to read: ", err)
	}
	if jsonObjTest.stringField != "string_value" {
		t.Errorf("Read invalid - stringField doesn't match. Expected: string_value | Got: %s", jsonObjTest.stringField)
	}
	if jsonObjTest.boolField != true {
		t.Errorf("Read invalid - boolField doesn't match. Expected: true | Got: %s", jsonObjTest.boolField)
	}
}
