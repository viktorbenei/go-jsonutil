package pathutil

import (
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

	if IsRelativePath("$ANENVVAR/some") {
		t.Error("$ANENVVAR/some should be absolute path!")
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
}
