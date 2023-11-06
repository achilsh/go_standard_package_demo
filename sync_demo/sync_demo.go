package syncdemo

import (
	"fmt"
	"sync"
)

// 本章节定义的是同步模块：
// 包括 sync.Once, sync.WaitGroup, sync.Pool, sync.Map, sync.Mutex, sync.Cond,

//sync.Once 保证多次调用时只有一次被执行； 可用于多协程下 业务只做一次逻辑， 也可以用于多次被调用时只有一次生效，比如单例创建对象时，只有一个对象被创建的场景。

type GeneralData interface {
	//proc1()
	//proc2()
	//....
}
type SingletonDemo struct{   //不管多少次获取内部数据，只获得一份值。 //singleton factory	
	once sync.Once  // An Once must not be copied after first use.
	//data 
	data string 
	generalObj  any // 或者是自定义接口：GeneralData ， 这是通用的数据模型， 可以用泛型或者用某些接口来替代。
	buildGeneralData func() any //这是产生数据的函数，统一使用一个函数类型来定义。
}

func (o *SingletonDemo) GetInstance() any {//对外提供接口，获取单例的对象。
	 o.once.Do(func() {
		o.generalObj = o.buildGeneralData() //获取对象， 多次调用，只能获取一次。
	 })
	 return o.generalObj
}


type IGetGeneralData interface { //那么定义 var xy IGetGeneralData = NewOnceDemo()
	GetInstance() any
}


func SyncOnceRun() {
	singletonMaps  := make(map[string]IGetGeneralData)
	singletonMaps["a"]=  NewOnceDemo()
	singletonMaps["b"] = NewOnceDemo2()
	for _, v := range singletonMaps {
		v.GetInstance()
		v.GetInstance()
	}
}

func NewOnceDemo() IGetGeneralData { //*SingletonDemo
	r := &SingletonDemo{
		buildGeneralData: func()any { //创建对象的具体实现。
			fmt.Println("call build general data obj=>int .")
			return  1
		},
	}
	return r
}
func NewOnceDemo2() IGetGeneralData { //*SingletonDemo
	r := &SingletonDemo{
				buildGeneralData: func() any {
					fmt.Println("call build general data obj=> string .")
					return "this test2."
				},
	}
	return r
}



func (o *SingletonDemo) init_call() {
	o.data = "run only once"
}
func (o *SingletonDemo) Get() string { //虽然 该方法可以多次被调用，但是对资源data的初始化只有一次。
	o.once.Do(o.init_call) //Do()的入参可以是 一个普通的函数，也可以一个struct的成员函数。
	return o.data
}
func (o *SingletonDemo) reset() {
	o.once = sync.Once{}
	o.data = "" 
}

//另外 sync.Once使用场景：定义创建某些资源的函数，比如：
// var (
//         mongoOnce sync.Once
//         mongoc    *mongo.Client
// )
// // MongoGet get mongo client
// func MongoGet() *mongo.Client {
//         mongoOnce.Do(func() {
//                 uri := fmt.Sprintf("mongodb://%s:27017/?retryWrites=false&directConnection=true", StoreHost)
//                 ctx := context.Background()
//                 logger.Infof("connecting to mongo: %s", uri)
//                 client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
//                 dtmimp.E2P(err)
//                 logger.Infof("connected to mongo: %s", uri)
//                 mongoc = client
//         })
//         return mongoc
// }
