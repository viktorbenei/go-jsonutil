package pathutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestIsRelativePath(t *testing.T) {
	t.Log("should return true if relative path, false if absolute path")

	if !IsRelativePath("./rel") {
		t.Error("./rel should be relative path!")
	}

	if IsRelativePath("/abs") {
		t.Error("/abs should be absolute path!")
	}

	if IsRelativePath("$THISENVDOESNTEXIST/some") {
		t.Error("$THISENVDOESNTEXIST/some should be absolute path! (any env)")
	}

	if !IsRelativePath("rel") {
		t.Error("'rel' should be relative path!")
	}
}

func TestIsPathExists(t *testing.T) {
	t.Log("should return false if path doesn't exist")

	exists, err := IsPathExists("this/should/not/exist")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if exists {
		t.Error("Should NOT exist")
	}

	exists, err = IsPathExists(".")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if !exists {
		t.Error("'.' Should exist")
	}

	exists, err = IsPathExists("")
	if err == nil {
		t.Error("Should return an error - no path provided!")
	}
	if exists {
		t.Error("'' (empty) Should NOT exist")
	}
}

func TestAbsPath(t *testing.T) {
	t.Log("should expand path")

	currDirPath, err := filepath.Abs(".")
	if err != nil {
		t.Error("Could not get current directory")
	}
	if currDirPath == "" || currDirPath == "." {
		t.Error("Invalid current dir path")
	}

	homePathEnv := "/path/home/test-user"
	err = os.Setenv("HOME", homePathEnv)
	if err != nil {
		t.Error("Could not set the ENV $HOME")
	}
	testFileRelPathFromHome := "some/file.ext"
	absPathToTestFile := fmt.Sprintf("%s/%s", homePathEnv, testFileRelPathFromHome)

	expandedPath, err := AbsPath("")
	if err == nil {
		t.Error("Should return an error for empty path")
	}
	if expandedPath != "" {
		t.Error("Empty path should be expanded to empty path. Got: ", expandedPath)
	}

	expandedPath, err = AbsPath(".")
	if err != nil {
		t.Error(err)
	}
	if expandedPath != currDirPath {
		t.Error("'.' Should be expanded to the current directory path. Got: ", expandedPath)
	}

	expandedPath, err = AbsPath(fmt.Sprintf("$HOME/%s", testFileRelPathFromHome))
	if err != nil {
		t.Error(err)
	}
	if expandedPath != absPathToTestFile {
		t.Error("Returned path doesn't match the expected path. :", expandedPath)
	}

	expandedPath, err = AbsPath(fmt.Sprintf("~/%s", testFileRelPathFromHome))
	if err != nil {
		t.Error(err)
	}
	if expandedPath != absPathToTestFile {
		t.Error("Returned path doesn't match the expected path. :", expandedPath)
	}
}
