package minit

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Reader interface for unit file reader
type Reader interface {
	// Next returns next field
	Next() (sec, key, val string, err error)
}

// Reader unit file stream reader
type reader struct {
	*bufio.Reader

	sec string // current section
}

// NewReader new unit file reader
func NewReader(r io.Reader) Reader {
	return &reader{
		Reader: bufio.NewReader(r),
	}
}

// Next returns next field
func (r *reader) Next() (sec string, key string, val string, err error) {
	var lineno int
	for {
		// increase line number
		lineno++

		// read line
		var line string
		if line, err = r.Reader.ReadString('\n'); err != nil {
			if err == io.EOF && len(line) > 0 {
				err = nil
			} else {
				return
			}
		}
		line = strings.TrimSpace(line)

		// ignore empty line
		if len(line) == 0 {
			continue
		}

		// ignore comment
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			r.sec = strings.TrimSpace(line[1 : len(line)-1])
			continue
		}

		// field
		splits := strings.SplitN(line, "=", 2)

		if len(splits) != 2 {
			err = fmt.Errorf("invalid syntax at line %d", lineno)
			return
		}

		sec = r.sec
		key = strings.TrimSpace(splits[0])
		val = strings.TrimSpace(splits[1])
		return
	}
}
