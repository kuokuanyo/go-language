package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = new(bytes.Buffer) //將bytes.Buffer指派給w
	_ = w.(*bytes.Buffer) //介面保存*bytes.Buffer
	_ = w.(io.ReadWriter) //*bytes.Buffer裡有Read and Writer method

	w = os.Stdout         //將os.Stdout指派給w
	_ = w.(*os.File)      //介面保存*os.File
	_ = w.(io.ReadCloser) //*os.File裡有Read and closer method
}
