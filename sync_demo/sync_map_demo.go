package syncdemo

import (
	"fmt"
	"reflect"
	"sync"
)

//sync.Map的特性： 多协程下安全； 在写少读多； 同时对不同的key上读写操作时。比 加锁和map的组合效率要高。

type SyncMapDemo struct {
	syncMap sync.Map 
}

func NewSyncMapDemo() *SyncMapDemo {
	r := &SyncMapDemo{}
	return r
}

func RunSyncMapDemo() {
	s := NewSyncMapDemo()
	fmt.Println("sync.Map.Store()")
	k := 123
	v := "abc"
	s.syncMap.Store(k, v)
	s.syncMap.Store(1, "a")
	s.syncMap.Store(2, 123)
	s.syncMap.Store(3, 912.12)
	s.syncMap.Store(123.123, 12323)
	

	fmt.Println("sync.Map.Load()")
	rv, i := s.syncMap.Load(k)
	if i {
		fmt.Println("load key: ", k, ", value: ", rv)
	} else {
		fmt.Println("load key: ", k, ", fail, not exist")
	}

	rv , i = s.syncMap.Load("no exist")
	if i {
		fmt.Println("load key: ", k, ", value: ", rv)
	} else {
		fmt.Println("load key: ", k, ", fail, not exist")
	}

	fmt.Println("sync.Map.Range()")
	s.syncMap.Range(func(k, v any) bool {
		rv := reflect.TypeOf(k)
		tt := rv.Kind()
		if  tt ==reflect.Int {
			fmt.Println("v : ", v)
			return true
		}
		if tt == reflect.Chan {
			fmt.Println("is chan type")
			return false
		}
		fmt.Println("tt: ", tt)
		return true
	})

	///...................
	s.syncMap.Range(func(k, v any) bool { // 在Range f方法内调用 sync.Map() 的其他方法。
		fmt.Println("delete k: ", k)
		s.syncMap.Delete(k)
		return true
	})

	fmt.Println("iter all keys in map: ")
	s.syncMap.Range(func(k, v any) bool {
		fmt.Println("k, ", k, ", v: ", v)
		return true
	})
}
