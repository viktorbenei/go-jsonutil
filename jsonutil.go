package jsonutil

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
// check out the _test to see some examples
//

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/viktorbenei/go-pathutil"
	"io"
	"os"
	"strings"
)

func ReadObjectFromJSONReader(reader io.Reader, v interface{}) error {
	jsonDecoder := json.NewDecoder(reader)
	if err := jsonDecoder.Decode(v); err != nil {
		return err
	}

	return nil
}

func ReadObjectFromJSONString(jsonString string, v interface{}) error {
	return ReadObjectFromJSONReader(strings.NewReader(jsonString), v)
}

func ReadObjectFromJSONFile(fpath string, v interface{}) error {
	file, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadObjectFromJSONReader(file, v)
}

func GenerateNonFormattedJSON(v interface{}) ([]byte, error) {
	jsonContBytes, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func GenerateFormattedJSON(v interface{}) ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func WriteObjectToJSONFile(fpath string, v interface{}) error {
	if fpath == "" {
		return errors.New("No path provided")
	}

	isExists, err := pathutil.IsPathExists(fpath)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New(fmt.Sprintf("File already exists at path: %s", fpath))
	}

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContBytes, err := GenerateFormattedJSON(v)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonContBytes)
	if err != nil {
		return err
	}

	return nil
}
