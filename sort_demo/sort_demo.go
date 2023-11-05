package sortdemo

import (
	"fmt"
	"sort"
)

/////该篇介绍： 分片和自定义容器的 排序。
//对自定义容器，比如各种类型的分片； 可以使用两种方式来排序：
// 方式 1： 使用 sort.Slice(分片，一个自定义的less函数对象)，下面实例会用到 .
// 方式2： 对任何类型的分片自定义一个类型，然后对实现该类型的 sort.Interface的方法，包括： Len() int， Less(i, j int) bool， Swap(i, j int)
//  然后调用 sort.Sort(自定义类型的分片）
type SortDemo struct {
}
func NewSortDemo() *SortDemo {
	return new(SortDemo)
}

func RunSortDemo() {
	sd := NewSortDemo()
	sd.sort_call()
	sd.sort_slice()
	sd.search()
}


type Person struct {
	Age int32
	Name string
}
func (d Person) String() string{
	return fmt.Sprintf("age: %v, name: %v", d.Age, d.Name)
}
type SPerson []Person 
func(s SPerson) Len() int {
	return len(s)
}

func (s SPerson) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] //直接上来就交换。
}

func (s SPerson) Less(i, j int) bool {
	t := []Person(s)
	if t[i] .Age < t[j].Age {
		return true
	}
	return false
}

func(p *SortDemo) sort_call() {
	fmt.Println("run sort.Sort()")
	data := []Person{
		{Age: 12, Name:"aa"}, {Age: 10, Name: "bb"},
	}
	sort.Sort(SPerson(data))
	for _, v := range data {
		fmt.Println(v)
	}
}

func(p *SortDemo)sort_slice() {
	fmt.Println("call sort.Slice()")

	data := []Person {
		{Age: 12, Name:"aa"}, {Age: 10, Name: "bb"},
	}
	sort.Slice(data, func(i, j int)bool { //提供一个自定义的Less函数。
		if data[i].Age < data[j].Age {
			return true
		}
		return false
	})
	//
	for _, v := range data {
		fmt.Println(v)
	}
}

func (p *SortDemo) search() {
	fmt.Println("run srot.Search()") //在有序的列表中，查找满足条件的元素在列表中的位置索引。 
	
	data := []Person {
		{Age: 12, Name:"aa"}, {Age: 10, Name: "bb"},
	}
	sort.Slice(data, func(i, j int)bool { //提供一个自定义的Less函数。
		if data[i].Age < data[j].Age {
			return true
		}
		return false
	})

	index := sort.Search(len(data), func(n int) bool {
		if data[n].Age > int32(10) {
			return true
		}
		return false
	})
	fmt.Println("sort.Search(),index: ", index)
}