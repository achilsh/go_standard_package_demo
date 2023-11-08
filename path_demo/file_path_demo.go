package pathdemo

import (
	"fmt"
	"os"
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

	fmt.Println("os.PathSeparator: ", string(os.PathSeparator), ", Separator: ", filepath.Separator)

	fmt.Println("filepath.ToSlash(): ", filepath.ToSlash("aa/bb/cc/"))
	fmt.Println("filepath.FromSlash(): ", filepath.FromSlash("aa/bb/cc/"))
	fmt.Println("filepath.Join(): ", filepath.Join("aa", "bb/cc", "/", "dd")) //将多个路径用 / 分隔符 连接成一个字符串。

	//Split拆分 紧跟在最后一个 / 之后的路径，将其分隔成目录和文件名组件; 不会删除 最后一个 /
	d, f :=  filepath.Split("x/y/z/")
	fmt.Printf("filepath.Split(): %v, %v\n",d, f)
	///返回带有路径的文件名的 文件类型。包括 . 
	fmt.Println("filepath: ", filepath.Ext("xx/y/z/xy.log"))

}
