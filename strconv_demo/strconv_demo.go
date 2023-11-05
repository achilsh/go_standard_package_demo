package strconvdemo

import (
	"fmt"
	"strconv"
)

type StrconvDemo struct {

}
func NewStrconvDemo() *StrconvDemo {
	return &StrconvDemo{}
}

func RunStrconvDemo() {
	sd := NewStrconvDemo()
	//
	sd.StringToBaseData()
	sd.BaseDataToString()

}

func (p *StrconvDemo) StringToBaseData() {
	fmt.Println("run string to base data() ")
	
	x, e:= strconv.Atoi("12321")
	fmt.Println("Atoi(): ", x, e)
	//
	b, e :=  strconv.ParseBool("false")
	fmt.Println("ParseBool(): ", b, e)
	b, e = strconv.ParseBool("true")
	fmt.Println("ParseBool(): ", b, e)

	// 2 for "0b", 8 for "0" or "0o",16 for "0x", 第二个参数的值范围，标识第一个字符串参数的格式。 
	// 0, 8,16, 32, 64: int, int8, int16, int32, and int64. 第三个参数的取值范围，标识返回值的类型。
	i, e := strconv.ParseUint("001011", 2 /**(0, 2,8, 10,16, to 36)*/, /**0, 8, 16, 32, and 64**/0 ) 
	fmt.Println("parseunint: ", i, e)

	i, e = strconv.ParseUint("011", 8 /**(0, 2,8, 10,16, to 36)*/, /**0, 8, 16, 32, and 64**/0 ) 
	fmt.Println("parseunint: ", i, e)

	i, e = strconv.ParseUint("011", 10 /**(0, 2,8, 10,16, to 36)*/, /**0, 8, 16, 32, and 64**/0 ) 
	fmt.Println("parseunint: ", i, e)

	i, e = strconv.ParseUint("a1f", 16 /**(0, 2,8, 10,16, to 36)*/, /**0, 8, 16, 32, and 64**/0 ) 
	fmt.Println("parseunint: ", i, e)

	//
	ii, e := strconv.ParseInt("01110", 2, 32)
	fmt.Println("parseInt(): ", ii, e)

	ii, e = strconv.ParseInt("17", 8, 32)
	fmt.Println("parseInt(): ", ii, e)
	
	ii, e = strconv.ParseInt("123", 10, 32)
	fmt.Println("parseInt(): ", ii, e)

	ii, e = strconv.ParseInt("123", 16, 32)
	fmt.Println("parseInt(): ", ii, e)
	//

	f, e := strconv.ParseFloat("123.123", 32)
	fmt.Println("parseFloat(): ", f, e)

	f, e = strconv.ParseFloat("123.12321", 64)
	fmt.Println("parseFloat(): ", f, e)

}

func(p *StrconvDemo) BaseDataToString() {
	fmt.Println("format bool: ")
	//bool 
	fmt.Println("bool str: ", strconv.FormatBool(false))
	fmt.Println("bool str: ", strconv.FormatBool(true))
	
	fmt.Println("format int: ")
	fmt.Println("int binary str: ", strconv.FormatInt(-100, 2))
	fmt.Println("int o str: ", strconv.FormatInt(-100, 8))
	fmt.Println("int d str: ", strconv.FormatInt(-100, 10))
	fmt.Println("int hex str: ", strconv.FormatInt(-100, 16))
	
	fmt.Println("format unit: ")
	fmt.Println("uint bin str: ", strconv.FormatUint(100,2))
	fmt.Println("uint o str: ", strconv.FormatUint(100,8))
	fmt.Println("uint d str: ", strconv.FormatUint(100,10))
	fmt.Println("uint hex str: ", strconv.FormatUint(100,16))

	fmt.Println("format float: ")
	fmt.Println("format float32: ", strconv.FormatFloat(123.123, 'f', -1 , 32))
	fmt.Println("float float64: ", strconv.FormatFloat(12311.3123, 'f', -1, 64))

	fmt.Println("strconv.Itoa(int)")
	fmt.Println(strconv.Itoa(123232))
}