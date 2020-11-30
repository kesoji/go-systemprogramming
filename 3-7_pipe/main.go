package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	header := bytes.NewBufferString("---- HEADER ----\n")
	contents := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("---- FOOTER ----\n")

	reader := io.MultiReader(header, contents, footer)
	io.Copy(os.Stdout, reader)

	var buffer bytes.Buffer
	reader2 := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader2, &buffer)
	// 読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	fmt.Println(buffer.String())
}
