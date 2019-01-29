package minit

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type TestLine struct {
	Sec string
	Key string
	Val string
}

var (
	TestUnitFileA = []string{
		" # this is a comment",
		"   ; this is a ini style comment  ",
		" key1 = val1",
		" [sec a  ]",
		" [ sec b  ]",
		" key2= val2 ",
		" key3 = val3 ",
		" [ sec c]",
		" key4 = val4 = val4 = val4 ",
	}
	TestUnitFileAContent = []TestLine{
		TestLine{Sec: "", Key: "key1", Val: "val1"},
		TestLine{Sec: "sec b", Key: "key2", Val: "val2"},
		TestLine{Sec: "sec b", Key: "key3", Val: "val3"},
		TestLine{Sec: "sec c", Key: "key4", Val: "val4 = val4 = val4"},
	}
)

func TestReader(t *testing.T) {
	r := NewReader(bytes.NewReader([]byte(strings.Join(TestUnitFileA, "\r\n"))))
	for i, l := range TestUnitFileAContent {
		sec, key, val, err := r.Next()
		if err != nil {
			t.Fatalf("Error: %s, item: %d", err.Error(), i)
		}
		if sec != l.Sec {
			t.Fatalf("bad sec: %s != %s, item: %d", sec, l.Sec, i)
		}
		if key != l.Key {
			t.Fatalf("bad key: %s != %s, item: %d", key, l.Key, i)
		}
		if val != l.Val {
			t.Fatalf("bad val: %s != %s, item: %d", val, l.Val, i)
		}
	}
	_, _, _, err := r.Next()
	if err != io.EOF {
		t.Fatalf("expect EOF, got %s", err.Error())
	}
}
