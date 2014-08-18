package jsonutil

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStructDecodingByPropertyVisibility(t *testing.T) {
	//
	// IMPORTANT:
	//  the JSON decoder will only decode struct attributes with Capital First Letter!
	//  so in case of:
	//		type MyType struct {
	//			ThisWill string
	//			thisWont string
	//		}
	//	only 'ThisWill' will get the proper value, 'thisWont' will be simply ignored!
	//

	testJsonString := `{
	"this_will": "it was decoded",
	"this_wont": "but this was not"
}`

	type MyType struct {
		ThisWill string `json:"this_will"`
		thisWont string `json:"this_wont"`
	}

	reader := strings.NewReader(testJsonString)
	jsonDecoder := json.NewDecoder(reader)
	var myObj MyType
	if err := jsonDecoder.Decode(&myObj); err != nil {
		t.Error("no error should occur, thisWont should be simply skipped. : ", err)
	}
	// Capital ThisWill will be decoded
	if myObj.ThisWill != "it was decoded" {
		t.Error("ThisWill was not correctly decoded (it should!)")
	}
	// thisWont WON'T be decoded, it will get it's default value.
	if myObj.thisWont != "" {
		t.Error("thisWont was decoded??")
	}

	//
	// Second issue - rather just a heads up for debugging:
	//	if you forget the ":" in the field comment it will once again simply skip the parameter
	// For example: `json:"something"` -> this is ok
	//				`json"something" -> this will be skipped!
	//
	type MySecondType struct {
		ThisWill string `json:"this_will"`
		ThisWont string `json"this_wont"`
	}
	reader = strings.NewReader(testJsonString)
	jsonDecoder = json.NewDecoder(reader)
	var myObj2 MySecondType
	if err := jsonDecoder.Decode(&myObj2); err != nil {
		t.Error("no error should occur, thisWont should be simply skipped. : ", err)
	}
	// Capital ThisWill will be decoded
	if myObj2.ThisWill != "it was decoded" {
		t.Error("ThisWill was not correctly decoded (it should!)")
	}
	// ThisWont WON'T be decoded, it will get it's default value because of the missing ":"
	if myObj2.ThisWont != "" {
		t.Error("ThisWont was decoded??")
	}

	//
	// And a correct, working solution:
	//
	type MyThirdType struct {
		ThisWill string `json:"this_will"`
		ThisToo  string `json:"this_wont"`
	}
	reader = strings.NewReader(testJsonString)
	jsonDecoder = json.NewDecoder(reader)
	var myObj3 MyThirdType
	if err := jsonDecoder.Decode(&myObj3); err != nil {
		t.Error("no error should occur, thisWont should be simply skipped. : ", err)
	}
	// Capital ThisWill will be decoded
	if myObj3.ThisWill != "it was decoded" {
		t.Error("ThisWill was not correctly decoded (it should!)")
	}
	// This time ThisToo is correct, so it will be decoded properly
	if myObj3.ThisToo != "but this was not" {
		t.Error("ThisToo was not correctly decoded (it should!)")
	}
}

type JsonObjTest struct {
	StringField string `json:"string_field"`
	BoolField   bool   `json:"bool_field"`
}

func TestGenerateNonFormattedJSON(t *testing.T) {
	testObj := JsonObjTest{
		StringField: "test string field",
		BoolField:   true,
	}
	jsonBytes, err := GenerateNonFormattedJSON(testObj)
	if err != nil {
		t.Error("Failed to generate non-formatted JSON: ", err)
	}

	expectedJSONString := `{"string_field":"test string field","bool_field":true}`
	if string(jsonBytes) != expectedJSONString {
		t.Error("Generated non formatted JSON doesn't match. Expected: ", expectedJSONString, " | Got: ", string(jsonBytes))
	}
}

func TestGenerateFormattedJSON(t *testing.T) {
	testObj := JsonObjTest{
		StringField: "test string field",
		BoolField:   true,
	}
	jsonBytes, err := GenerateFormattedJSON(testObj)
	if err != nil {
		t.Error("Failed to generate formatted JSON: ", err)
	}

	expectedJSONString := `{
	"string_field": "test string field",
	"bool_field": true
}`
	if string(jsonBytes) != expectedJSONString {
		t.Error("Generated formatted JSON doesn't match. Expected: ", expectedJSONString, " | Got: ", string(jsonBytes))
	}
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
	if jsonObjTest.StringField != "string_value" {
		t.Errorf("Read invalid - stringField doesn't match. Expected: string_value | Got: %s", jsonObjTest.StringField)
	}
	if jsonObjTest.BoolField != true {
		t.Errorf("Read invalid - boolField doesn't match. Expected: true | Got: %s", jsonObjTest.BoolField)
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
	if jsonObjTest.StringField != "string_value" {
		t.Errorf("Read invalid - stringField doesn't match. Expected: string_value | Got: %s", jsonObjTest.StringField)
	}
	if jsonObjTest.BoolField != true {
		t.Errorf("Read invalid - boolField doesn't match. Expected: true | Got: %s", jsonObjTest.BoolField)
	}
}
