package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestWriterReader(t *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	fmt.Fprintf(&b, "World")

	b.WriteTo(os.Stdout)
}
