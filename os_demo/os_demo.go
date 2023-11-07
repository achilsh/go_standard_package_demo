package osdemo

import (
	"fmt"
	"os"
)

// os包 提供了 访问 operating-system 操作系统 且与平台无关的接口。
// 这些接口返回的错误不是错误码数字，而是一些错误类型的对象。可以被 unpacked 。
//

func RunOsDeo() {
	fmt.Println("os pkg demo run:")
	fmt.Println("---- os constants -----")
	fmt.Println("os.O_RDONLY: ", os.O_RDONLY, os.O_WRONLY, os.O_RDWR, //只读打开文件，只写打开文件，读写文件 时指定的fLag，
		os.O_APPEND, os.O_CREATE, // 写文件时往末尾追加； 打开文件时文件不存在则创建。
	)
	fmt.Printf("%o, %s, %o, %o\n", os.ModeDir, os.DevNull, os.ModePerm, os.ModeType)

	fmt.Println("..... os variable ......")
	fmt.Println("os.Stdin: ", os.Stdin.Name(), ", os.Stdout: ", os.Stdout.Name(), ", os.Stderr: ", os.Stderr.Name())
	fmt.Println("os.Args: ", os.Args) //命令行参数，包括 运行的二进制名 和 参数值。

	fmt.Println("chmod(): ", os.Chmod("./x1.log", os.ModePerm |0645))

	fmt.Println("os.Chdir: ", os.Chdir("../")) //改变当前工作目录
	d, _ := os.Getwd()  //获取当前工作目录
	fmt.Println("new work dir: ", d)

	os.Chdir("./go_std_pkg_demo")
	fmt.Println("chmod(): ", os.Chmod("./x1.log", os.ModePerm & 0645))
	//
	fmt.Println("---------- os.Env: ---------")
	for _, v := range os.Environ() { //获取所有的环境变量。
		fmt.Println("env: ", v)
	}
	fmt.Println("os.Setenv(): ", os.Setenv("test_demo", "12312"))
	fmt.Println("os.Getenv(): ", os.Getenv("test_demo"))
	v, ok := os.LookupEnv("test_demo") //检索环境变量。
	if ok {
		fmt.Println("find env: test_demo, value: ", v) //环境变量存在，即使环境变量的值为空，也返回为true.
	} else {
		fmt.Println("not find env: test_demo value")
	}

	fmt.Println("os.UnSetenv(): ", os.Unsetenv("test_demo"))
	fmt.Println("os.Getenv(): ", os.Getenv("test_demo"))
	v, ok = os.LookupEnv("test_demo")
	if ok {
		fmt.Println("find env: test_demo, value: ", v)
	} else {
		fmt.Println("not find env: test_demo value")
	}

	fmt.Println("-- run start path: ---")
	p, _ := os.Executable() //返回启动当前进程的可执行文件的路径名
	fmt.Println("path: ", p)

	fmt.Println("exit process......")
	// os.Exit(1) // 以指定的错误码退出进程。
	fmt.Println("exit process end......")


	fmt.Println("----create hard link") //给文件创建一个硬链接。 
	os.Link("x1.log", "xxxx1.log")

	fmt.Println("----symbolic link")
	os.Symlink("x1.log", "xxxx2.log")

	fmt.Println(".......create dir with perm .......")
	os.Mkdir("xxx3_dir", 0750) 
	
	fmt.Println("...... create any parent dir.........")
	os.MkdirAll("./xxx4/yyyy1", 0750)


	
}
