package osdemo

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// os包 提供了 访问与平台无关的  operating-system 接口。与系统相关的系统接口的包是： syscall
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

	fmt.Println("chmod(): ", os.Chmod("./x1.log", os.ModePerm|0645))

	fmt.Println("os.Chdir: ", os.Chdir("../")) //改变当前工作目录
	d, _ := os.Getwd()                         //获取当前工作目录
	fmt.Println("new work dir: ", d)

	os.Chdir("./go_std_pkg_demo")
	fmt.Println("chmod(): ", os.Chmod("./x1.log", os.ModePerm&0645))
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
	v, ok = os.LookupEnv("test_demo") //查找环境变量
	if ok {
		fmt.Println("find env: test_demo, value: ", v) // 环境变量存在，即使变量值为空也返回true.
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

	fmt.Println("----symbolic link") //创建文件的软连接。
	os.Symlink("x1.log", "xxxx2.log")

	fmt.Println(".......create dir with perm .......")
	os.Mkdir("xxx3_dir", 0750) //创建目录

	fmt.Println("...... create any parent dir.........")
	os.MkdirAll("./xxx4/yyyy1", 0750) //创建多级目录

	fmt.Println("----- file write, read, rename, remove, ")
	os.WriteFile("./x2.log", []byte("thi is write file demo"), 0666) //文件不存在，则以0666权限去创建该文件。
	readRet, e := os.ReadFile("x2.log")                              //返回错误 e 标识读失败。 io.EOF 不会作为错误，就是说 返回的错误中没有 io.EOF
	if e != nil {
		fmt.Println("read data fail, e: ", e)
	} else {
		fmt.Println("read file, file data: ", string(readRet))
	}
	e = os.Remove("x1.log")
	if e != nil {
		fmt.Println("remove x1.log fail, e: ", e)
	}
	e = os.Remove("x2.log") //删除文件或者目录
	if e != nil {
		fmt.Println("remove x2.log fail, e: ", e)
	}
	e = os.Remove("xxx3_dir")
	if e != nil {
		fmt.Println("remove xxx3_dir fail, e: ", e)
	}
	e = os.Remove("xxxx2.log")
	if e != nil {
		fmt.Println("remove xxxx2.log fail, e: ", e)
	}
	e = os.Remove("xxx4")
	if e != nil {
		fmt.Println("remove xxx4 fail, e: ", e)
	}
	os.Remove("xxxx1.log")

	DirOpRun() //

	FileRun()
	FileInfo()

	ProcessRun()

	SignalDemo()
}

func DirOpRun() {
	fmt.Println("--- read dir info: ")
	ds, e := os.ReadDir(".") // // 读取指定目录的信息。
	if e != nil {
		fmt.Println("read cur dir fail, e: ", e)
	} else {
		for _, d := range ds {
			f, _ := d.Info()
			fmt.Println("name: ", d.Name(), ", is dir: ", d.IsDir(), ", type: ", d.Type(), ", info: ", f)
		}
	}
	// file.ReadDir()
}

func FileRun() {
	fmt.Println("-------- file ops ------")
	f, e := os.Create("new_file_or_truncate_old_file.log") //创建一个新文件，或者切断一个老文件. 可读写该创建的 文件。
	if e != nil {
		fmt.Println("op create fail, e: ", e)
	} else {
		fmt.Println("succ: you can write and read on this created file. file name: ", f.Name())
		//
		_ = f // 可以对该文件进行读写操作。
	}
	f.Close() //需要关闭打开的文件。

	fmt.Println("----- open file within only read mode.")
	f1, e := os.Open("new_file_or_truncate_old_file.log") //打开一个文件，用于只能读操作。
	if e != nil {
		fmt.Println("open file for only read fail, e: ", e)
	} else {
		_ = f1
		fmt.Println("open file for only read succ. name: ", f1.Name())
	}
	f1.Close() //需要关闭打开的文件。

	fmt.Println("---- open file with self defined w/r/create and perm")
	f2, e := os.OpenFile("xyz.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664) //指定模式，权限 打开一个文件。
	if e != nil {
		fmt.Println("open file fail, e: ", e)
	} else {
		_ = f2
		fmt.Println("OpenFile succ. fileName: ", f2.Name())
		fmt.Println("write data to file: ", f2.Name())
		f2.Write([]byte("1) this is demo write file..")) //写文件，因为是 append 方式创建的文件，所以写文件是一直追加在尾部。
		f2.Write([]byte("2) this is demo write file.."))
		///
		fmt.Println("read data from file: ", f2.Name())
		dst := make([]byte, 4)
		n, e := f2.ReadAt(dst, 0) // 如果直接 调用 f2.Read()， 是直接返回在io.EOF,因为f2确实也是 append  write 操作。
		if e != nil {
			fmt.Println("read from fail, e: ", e)
		} else {
			fmt.Println("read from file, succ, data: ", string(dst), ", len: ", n)
		}

		//
		fmt.Println("print file info: ")
		f2Info, e := f2.Stat()
		if e != nil {
			fmt.Println("get f2 stat fail, e: ", e)
		} else {
			fmt.Println("f2 stat info: ", f2Info.Name(), f2Info.IsDir(), f2Info.Size(), f2Info.Mode())
		}
	}
	f2.Close()

	//
	fmt.Println("--- open dir and then call Chdir() to this dir")
	f3, e := os.Open(".")
	if e != nil {
		fmt.Println("open dir fail, e: ", e)
	} else {
		fmt.Println("open dir succ, dir name:", f3.Name())
		fmt.Println("change cur working dir to: ", f3.Name(), f3.Chdir())
	}

}

func FileInfo() {
	fmt.Println("get file info by file name:")
	fi, e := os.Stat("xyz.log")
	if e != nil {
		fmt.Println("get file info fail for : ", "xyz.log")
	} else {
		fmt.Printf("file info: %v, %d, %d, %v, %v\n", fi.Name(), fi.Mode().Perm(), fi.Mode().Type(), fi.IsDir(), fi.Sys())
	}
}

func ProcessRun() {
	//exec 包主要作用： 执行一些外部命令。 
	fmt.Println("run exec.Command to run external command...")
	//这是进程相关的操作，建议使用 exec 包中接口。 包括: exec runs 外部 commands
	cmd := exec.Command("ls", "-lt", "./") //返回一个执行命令的cmd 结构体。
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr 

	e := cmd.Run()      //执行并等待命令结束。
	if e != nil {
		fmt.Println("run external cmd fail, e: ", e)
	}

	fmt.Println("----------- start and run: ")
	cmd2 := exec.Command("ls", "-lt", "./")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr 
	e = cmd2.Start()  //运行外部命令，但是不等待命令运行结束
	if e != nil {
		fmt.Println("start cmd fail, e: ", e)
	}

	cmd2.Wait() //等待命令的结束，包括输入输出的copy 完成。
}

func SignalDemo() {
	//忽略某些信号， 比如：
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM) //忽略 syscall.SIGINT 信号。就不会被捕获到。

	chSig := make(chan os.Signal) //定义os.Signal 量。
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)    // 将接收到的信号发给 chan
	fmt.Println("wait to ctrl +c  stop the cmd")
	<-chSig
	// syscall 包的主要作用：Package syscall contains an interface to the low-level operating system primitives. 
	// The details vary depending on the underlying system, and by default, 
	// godoc will display the syscall documentation for the current system

	//signal 包的主要作用： 实现了对输入信号的访问。包括： Notify, Ignore等操作。
}
