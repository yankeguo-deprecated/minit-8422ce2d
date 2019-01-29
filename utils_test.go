package minit

import (
	"io/ioutil"
	"os"
	"testing"
)

const TestFile = "/tmp/test-service.service"

func TestSearchUnitFile(t *testing.T) {
	var ret string
	var err error

	os.Remove(TestFile)
	if ret, err = SearchUnitFile("test-service.service"); err == nil {
		t.Fatal("should failed")
	}
	SearchPaths = append(SearchPaths, "/tmp")
	if ret, err = SearchUnitFile("test-service.service"); err == nil {
		t.Fatal("should failed")
	}
	if err = ioutil.WriteFile(TestFile, []byte("dummy=true"), 0644); err != nil {
		t.Fatalf("failed to write: %s", err.Error())
	}
	if ret, err = SearchUnitFile("test-service.service"); err != nil {
		t.Fatalf("should success: %s", err.Error())
	}

	if ret != TestFile {
		t.Fatalf("should equal: %s != %s", ret, TestFile)
	}
}
