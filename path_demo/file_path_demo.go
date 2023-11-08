package pathdemo

import (
	"fmt"
	"path/filepath"
)

//filepath包 主要是用来操作操作系统上适合的文件路径。

func RunFilePathDemo() {
	fmt.Println("---------- run filepath demo()------------")
	var (
		f string 
	 	e error
	)
	///////////// filepath.Abs()
	f, e = filepath.Abs("abc.log") //获取当前文件的绝对路径，等于当前工作目录 和 文件名组合成完整路径。
	fmt.Printf("absoluted path for file: %v, %v\n", f, e)

	//// 文件路径的最后一段。 filepath.Base()
	fmt.Println("run file name last item: ", filepath.Base("a/b/c.log"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))


}

