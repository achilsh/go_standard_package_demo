package pathdemo

import (
	"fmt"
	"path"
)

//path包的主要作用： 操作 斜杠分割符路径。 就是仅用于 斜杠分割符 分割的路径。比如 url. 路径等。
//如果要对操作系统文件路径进行操作，可以使用 path/filepath 包。

func RunPathDemo() {
	fmt.Println("---- run path package to op path of slashes.")
	//获取路径最后一段 path.Base()
	fmt.Println("get last item of path: ", path.Base("/a/b/c") )  ///返回 c
	fmt.Println("get last itme of path: ", path.Base("/a/b/c/"))  //返回 c
	fmt.Println("get last itme of path: ", path.Base("/"))      //返回  /
	fmt.Println("get last item of path: ", path.Base(""))      //返回 .

	//获取路径目录 path.Dir()
	fmt.Println("get except last item of path: ", path.Dir("/a/b/c")) //返回  /a/b
	fmt.Println("get except last item of path: ", path.Dir("a/b/c"))  //返回  a/b
	fmt.Println("get except last item of path: ", path.Dir("/a/"))    //返回  /a
	fmt.Println("get except last item of path: ", path.Dir("a/"))     //返回  a
	fmt.Println("get except last item of path: ", path.Dir("/"))      //返回  /
	fmt.Println("get except last item of path: ", path.Dir(""))       //返回  .

	//获取文件名的后续类型。 path.Ext()
	fmt.Println("get file name extension: ", path.Ext("a/b/c/e.log")) 	//返回 .log
	fmt.Println("get file name extension: ", path.Ext("a/b/c/e")) 		//返回 空
	fmt.Println("get file name extension: ", path.Ext(""))          	//返回 空

	//判断路径是否 是绝对路径 . path.IsAbs()
	fmt.Println("check path is 绝对 path: ", path.IsAbs("a/b/c"))  //返回 false
	fmt.Println("check path is 绝对 path: ", path.IsAbs("/a/b/c")) //返回 true
	fmt.Println("check path is 绝对 path: ", path.IsAbs("/"))      //返回 true 
	fmt.Println("check path is 绝对 path: ", path.IsAbs(""))       //返回 false

	//将多个路径 items 用斜杠连接成一条路径。 path.Join()
	fmt.Println("contect some path item into a whole path: ", path.Join("a", "b")) //返回 a/b
	fmt.Println("contect some path item into a whole path: ", path.Join("a", "b/c")) //返回 a/b/c
	fmt.Println("contect some path item into a whole path: ", path.Join("a", "../b/c")) //返回  b/c
	fmt.Println("contect some path item into a whole path: ", path.Join("a", "")) //返回  a
	fmt.Println("contect some path item into a whole path: ", path.Join("", "a")) //返回  a


	//将路径 划分为 目录和 文件名， 适合分离目录和文件名。 base.Split()
	var d, f string 
	d, f =  path.Split("a/b/c.log")
	fmt.Printf("split path into dir and file name, dir: %v, filename: %v\n", d,f ) //返回： a/b/  和 c.log

	d, f =  path.Split("c.log")
	fmt.Printf("split path into dir and file name, dir: %v, filename: %v\n", d,f ) //返回：  和 c.log

	d, f =  path.Split("c")
	fmt.Printf("split path into dir and file name, dir: %v, filename: %v\n", d,f )  //返回：   和 c

	d, f =  path.Split("/")
	fmt.Printf("split path into dir and file name, dir: %v, filename: %v\n", d,f ) //返回：  / 和 

	d, f =  path.Split("")
	fmt.Printf("split path into dir and file name, dir: %v, filename: %v\n", d,f ) //返回：   和  
}