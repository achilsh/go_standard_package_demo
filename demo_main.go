package main

import (
	"std_pkg_demo/bufio_demo"
	"std_pkg_demo/builtin_demo"
	"std_pkg_demo/bytes_demo"
)

func main() {
	bufio_demo.Bufio_demo_run()
	builtin_demo.BuiltinDemoRun()
	bytes_demo.BytesDemoRun()

	//
	bytes_demo.BytesBufferDemo()
	//
	bytes_demo.BytesReaderRun()
}
