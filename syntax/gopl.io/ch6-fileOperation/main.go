package main

import (
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch6-fileOperation/demo1"
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch6-fileOperation/demo2-IO"
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch6-fileOperation/demo3-output"
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch6-fileOperation/demo4-fileCopy"
)

func main() {

	demo1.OpenFile()
	demo2io.IoFunc() // ASCII码 需要手动转换
	demo2io.IOBuffer()
	demo3output.Test()
	demo4filecopy.Test()
}
