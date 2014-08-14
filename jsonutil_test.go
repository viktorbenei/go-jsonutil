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
}
