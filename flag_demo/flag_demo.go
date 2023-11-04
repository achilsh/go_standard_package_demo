package flag_demo

import (
	"flag"
	"fmt"
)

//实现了： 命令行标记 的解析。
//简单实用flag 包，一般是定义一些flag 变量，调用： flag.Parse()即可。
//有两种方式来使用：
//1. 通过 var busiFlag = flag.Xxx() 返回flag变量指针。后续就可以使用这些返回的变量。
//2. 自定义的flag 变量。然后 传入函数中 flag.XxxVar()。后续就可以使用这些定义的变量。

//方式1：
var (
	// 函数入参解释：        flag 名字， flag默认值， flag的业务含义
	BoolFlagPtr = flag.Bool("boolFlag", true, "explain for this  bool flag")
	IntFlagPtr = flag.Int("iFlag", 1, "explain for this int flag")
	UintFlagPtr = flag.Uint("uiFlag", 2, "explain for this uint flag")
	Int64FlagPtr = flag.Int64("int64Flag", 3, "explain for this int64 flag")
	Uint64FlagPtr = flag.Uint64("ui64Flag", 4, "explain for this uint64 flag")
	Float64FlagPtr = flag.Float64("float64Flag", 5.0, "explain for this float64 flag")
	StrFlagPtr = flag.String("strFlag", "6", "explain for this string flag")
	//指定的flag 名字 可以在 编译成的二进制文件 --help 打印出来。
	// -flag // 只支持 bool 类型
	// -flag=x
	// -flag x // 只支持非 bool 类型
)

//方式2： 定义flag 变量。
var (
	BoolFlag bool 
	IntFlag int 
	UintFlag uint 
	Int64Flag int64 
	Uint64Flag uint64
	FLoat64Flag float64 
	StrFlag string 
)

func init () {
	// 函数的入参包括： 已经定义的flag变量名， flag名字， 默认值， flag的业务含义
	flag.BoolVar(&BoolFlag, "boolFlagEx",  true , "explain for int bool flag ex")
	flag.IntVar(&IntFlag, "intFlagEx", 11, "explain for int flag ex") 
	flag.UintVar(&UintFlag, "uintFlagEx", 22, "explain for uint flag ex")
	flag.Int64Var(&Int64Flag, "int64FlagEx", 33, "explain for int64 ex")
	flag.Uint64Var(&Uint64Flag, "uint64FlagEx", 44, "explain for uint64 ex")
	flag.Float64Var(&FLoat64Flag, "float64FlagEx", 55.55, "explain for float64 ex")
	flag.StringVar(&StrFlag, "strFlagEx", "66", "explain for string ex") 
}
 


type FlagDemo struct {
	//定义flag变量 ,或者定义在全局变量中并进行初始化。
}

func NewFlagDemo() *FlagDemo {
	r := &FlagDemo{}
	
	flag.Parse() //调用函数用于解析flag.
	flag.Usage() //用于当命令行参数输入有问题时，可以在该判断条件下 运行该语句。
	return r
}

func (p *FlagDemo) printBoolFlag() {
	fmt.Println("bool: ", *BoolFlagPtr, ", ", BoolFlag)
}
func (p *FlagDemo) printIntFlag() {
	fmt.Println("int: ", *IntFlagPtr, ", ", IntFlag)
}
func (P *FlagDemo) printUintFlag() {
	fmt.Println("uint: ", *UintFlagPtr, ", ", UintFlag)
}
func (p *FlagDemo) printInt64() {
	fmt.Println("int64: ", *Int64FlagPtr, ", ", Int64Flag)
}
func (p *FlagDemo) printUint64() {
	fmt.Println("uint64: ", *Uint64FlagPtr, ", ", Uint64Flag)
}
func(p *FlagDemo) printFloat64() {
	fmt.Println("float64: ", *Float64FlagPtr, ", ", FLoat64Flag)
}
func (p *FlagDemo) printString() {
	fmt.Println("string: ", *StrFlagPtr,", ", StrFlag)
}

func RunFlagDemo() {
	fmt.Println(">>>> begin to run flag demo >>>>>")
	 r := NewFlagDemo() 
	 //
	 r.printBoolFlag()
	 r.printIntFlag()
	 r.printUintFlag()
	 r.printInt64()
	 r.printUint64()
	 r.printFloat64()
	 r.printString()
}
