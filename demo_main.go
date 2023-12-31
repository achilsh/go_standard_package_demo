package main

import (
	"std_pkg_demo/bufio_demo"
	"std_pkg_demo/builtin_demo"
	"std_pkg_demo/bytes_demo"
	"std_pkg_demo/encode_demo"
	"std_pkg_demo/flag_demo"
	genericdemo "std_pkg_demo/generics_demo"
	iodemo "std_pkg_demo/io_demo"
	osdemo "std_pkg_demo/os_demo"
	pathdemo "std_pkg_demo/path_demo"
	runtimedemo "std_pkg_demo/runtime_demo"
	sliceofanydemo "std_pkg_demo/slice_of_any_demo"
	sortdemo "std_pkg_demo/sort_demo"
	strconvdemo "std_pkg_demo/strconv_demo"
	"std_pkg_demo/strings_demo"
	syncdemo "std_pkg_demo/sync_demo"
	templateDemo "std_pkg_demo/template"
	timedemo "std_pkg_demo/time_demo"
)

func main() {
	bufio_demo.Bufio_demo_run()
	builtin_demo.BuiltinDemoRun()
	bytes_demo.BytesDemoRun()

	//
	bytes_demo.BytesBufferDemo()
	//
	bytes_demo.BytesReaderRun()
	//
	encode_demo.RunEncodeDemo()
	encode_demo.RunBinaryDemo()
	//
	flag_demo.RunFlagDemo()
	//
	strings_demo.StringsDemoRun()
	strings_demo.RunStringBuilderDemo()
	strings_demo.RunStringsReader()
	//
	strconvdemo.RunStrconvDemo()
	//
	sliceofanydemo.RunSliceOfAny()
	//
	sortdemo.RunSortDemo()
	//
	timedemo.RunTimeDemo()

	syncdemo.SyncOnceRun()
	syncdemo.RunGenericSingleton()
	syncdemo.RunSyncPoolDemo()
	syncdemo.RunBytesBuffPool()
	//
	syncdemo.RunSyncMapDemo()
	syncdemo.RunAtomicDemo()
	//
	runtimedemo.RunDemo()
	//
	osdemo.RunOsDeo()
	pathdemo.RunPathDemo()
	pathdemo.RunFilePathDemo()
	//
	iodemo.RunIoPackageDemo()
	iodemo.IoTypeDemoRun()
	//
	templateDemo.TemplateDemoRun()
	templateDemo.RunGenStructType()
	//
	genericdemo.RunGenerics()
}
