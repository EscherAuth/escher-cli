package reverseproxy

import (
	"strings"
)

type myReader struct {
	s *strings.Reader
}

func (m *myReader) Close() error {
	return nil
}

func (m *myReader) Read(p []byte) (n int, err error) {
	return m.s.Read(p)
}

func newMyReader(s string) *myReader {
	return &myReader{s: strings.NewReader(s)}
}
