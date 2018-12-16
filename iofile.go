package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func IOExample() {

	var b = bytes.Buffer{}

	b.Write([]byte("Hello"))

	fmt.Fprintf(&b,"World")

	io.Copy(os.Stdout,&b)
}

