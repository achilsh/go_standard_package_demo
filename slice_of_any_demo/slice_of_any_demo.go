package sliceofanydemo

import (
	"cmp"
	"fmt"
	"slices"
)

// 该篇主要介绍：存储任何类型的分片 及其这些分片的操作。
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：
// 任何数据类型的分片及其函数：

type SliceOfAnyDemo struct {

}

func NewSliceOfAnyDemo() *SliceOfAnyDemo {
	return new(SliceOfAnyDemo)
}

func RunSliceOfAny() {
	sa := NewSliceOfAnyDemo()
	sa.binarySearch()

	sa.sortFunc()

}

//在任何类型中分片内做二分查找，基础条件是： 该分片是有序的。
func (p *SliceOfAnyDemo) binarySearch() {
	fmt.Println("slices.BinarySearch()")
	xy := []string{"ao3kd", "bi123o", "ci12", "d123"}
	i, e := slices.BinarySearch(xy, "ci12")         //在有序列表中查找一个值，存在则返回值的索引位置。
	fmt.Println("binary seach by slices, find: ", i, e)
	//

	fmt.Println("slices.BinarySearchFunc()")
	xy = []string{"ao3kd", "bi123o", "ci12", "d123"}
	
	//在有序列表（递增的）中，使用自定义函数查找一个元素, 如果该元素和分片中的元素相等。存在则返回索引。
	//其中自定义函数就是一个比较函数： 元素相等则返回0， 分片内元素小于该元素，则返回-1，  分片元素大于该元素，怎返回1.
	//该分片必须是递增有序的。
	//适用于自定义类型的分片。在该类型分片中，使用特定的比较函数查找元素在分片中的位置。
	r, b :=  slices.BinarySearchFunc(xy, "ci12",  func(sliceItem, targetItem string)int {
		return cmp.Compare(sliceItem, targetItem)
	})
	fmt.Println("BinarySearchFunc: ", r, b)
}


func (p *SliceOfAnyDemo) sortFunc() {
	fmt.Println("run slices.SortFunc()")
	//排序时，指定特定的比较函数。 特别适用于 对自定义类型的 分片排序。排序的比较函数也是可以自定义的。
	xy := []string{"aa", "bb", "123", "ac"}
	slices.SortFunc(xy, func(a, b string) int{
		return cmp.Compare(a,b)
	})

	//稳定排序
	slices.SortStableFunc(xy, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	//标准的排序：
	slices.Sort(xy)
}

func (p *SliceOfAnyDemo) reverse() {
	//反转。
}

// 
func (p *SliceOfAnyDemo) contains() {
	//分片中是佛包含某个元素：
	xy := []string{"aa", "bb", "cc"}
	slices.Contains(xy, "bb")

	//判断分片中元素是否满足特定函数，如果满足函数返回true .
	slices.ContainsFunc(xy, func(a string) bool {
		if len(a) <= 0 {
			return true
		}
		return false
	})
}


func (p *SliceOfAnyDemo) Index() {
	//....在分片中查找，某个元素第一出现的位置：
	xy := []string{"aaa", "bbb", "ccc"}
	slices.Index(xy, "aaa")
	//

	//在分片中元素，查找满足特性（自定义函数的）的第一个元素：
	slices.IndexFunc(xy, func(d string) bool {
		if len(xy) <= 0 {
			return true
		}
		return false
	})
}

