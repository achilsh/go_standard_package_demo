package timedemo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//本章节介绍 时间相关的点：
// 数字转化为时间：  将具体数字强转为 timee.Duration， time.Duration(10).  然后再乘以 具体的单位， 比如：time.Second; time.Millisecond, time.Microsecond 等。

//layout: 用于解析或者格式化时间。

type TimeDemo struct {
}

func NewTimeDemo() *TimeDemo {
	return new(TimeDemo)
}

func RunTimeDemo() {
	t := NewTimeDemo()
	//....
	t.after()
	t.tick()
	t.parseDuration()
	t.since()
	t.hour()
	t.string_call()
	t.local()
	t.ticker()
	//
	t.time_struct_func()
}

func (t *TimeDemo) after() { //从现在 多少时间之后的点。
	fmt.Println("call time.After(), ", time.Now().Unix())
	//等待一段时间后向 返回的chan发送一个通知； 调用方 阻塞等待chan通知。可以用来等待将来的某个时间点。
	c := time.After(2 * time.Second) //下一个 n time的时间点。
	select {
	case <-c:
		fmt.Println("等待一段时间后，当前时间已经到了。") //可用于超时计时器。
	}
	fmt.Println("end time: ", time.Now().Unix())

	// sleep:
	fmt.Println("time.Sleep(1*time.Second)")
	time.Sleep(1 * time.Second)

}

func (t *TimeDemo) tick() {
	//主要包装了 time.NewTicker ,
	fmt.Println("run time.Tick()")
	tt := time.Tick(1 * time.Second) //返回一个chan, 系统会定时往chan 发送一个信号。 调用方 可以不断的消费chan 实现定时器。

	i := 0
	for range tt {
		fmt.Println("recv tick: ", time.Now().Unix())
		if i > 1 {
			fmt.Println("begin to break tick.")
			break
		}
		i++
	}
}
func (t *TimeDemo) parseDuration() {
	fmt.Println("run time.ParseDuration()")
	h, _ := time.ParseDuration("1h") // parses a duration string.  主要用来将字符串的时间 转化 time.Duration. 其中：字符串中出现的单位有： "ns", "us" (or "µs"), "ms", "s", "m", "h"
	fmt.Printf("hour: %#v\n", h/time.Hour)
	s, _ := time.ParseDuration("3s")
	fmt.Printf("second: %#v\n", s/time.Second)

}

func (t *TimeDemo) since() {
	fmt.Println("time.Since(t)") //主要用来计算之前的某个时间点 到现在时刻的 时间 time.Duration .通常是用于计算时间差。 参考点是过去和当前time.Now().
	before := time.Now()
	<-time.After(2 * time.Second)
	elapsed := time.Since(before) //计算过去到现在的时间差。
	fmt.Printf("%v\n", elapsed)
}

func (t *TimeDemo) hour() {
	fmt.Println("time.Duration.Hour()") //返回指定时间的一共多少hour。
	d, _ := time.ParseDuration("2h")

	fmt.Println("hour: ", d.Hours())              //time.Duration 一共折算成多少 hour
	fmt.Println("minute: ", d.Minutes())          //time.Duration 一共 折算成多少 分钟
	fmt.Println("second: ", d.Seconds())          //time.Duration 一共折算成多少 秒
	fmt.Println("millSecond: ", d.Milliseconds()) //time.Duration 一共折算成多少 毫秒

	d2, _ := time.ParseDuration("1h15m30.918273645s")
	fmt.Println(d2.Hours(), d2.Minutes(), d2.Seconds(), d2.String())
}

func (t *TimeDemo) string_call() {
	fmt.Println("time.Duration.String()") //将 time.Duration 的值转化为 字符串。比如： 2*time.Hour 可以转化 2h ，他的反操作是time.ParseDuration()
	a := time.Hour*1 + time.Minute*2 + time.Second*3 + time.Millisecond*4 + time.Microsecond*5
	fmt.Println("time.Duration.String() ret: ", fmt.Sprintf("%s", a))
}

func (t *TimeDemo) local() {
	//有个全局的指针变量 time.Local, 标识 本地 time.Location 标识本地的time zone.
	fmt.Println("time.Location	")
	fmt.Println(time.Local.String())

	fmt.Println("run time.LoadLocation()")
	time.LoadLocation("Local") //获取本地的location
	time.Now().Location()
}

func (t *TimeDemo) ticker() {
	fmt.Println("time.Ticker")            //定时器，  他内部保存了一个chan, 系统会定期往 chan 发送 tick. 调用方可 不断的从 chan 读取tick 实现定时。
	tk := time.NewTicker(1 * time.Second) //创建一个指定定时间隔的 ticker.
	defer tk.Stop()                       //停掉 该定时器。

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	i := 0
	for {
		select {
		case <-tk.C: //此时接收到定时器 通知。 因为系统往定时器的chan 发送了tick。
			fmt.Println("tick arrive, you can press ctrl +c stop.")
			i++
			if i == 5 {
				fmt.Println("call Reset(d): ")
				tk.Reset(2 * time.Second) // 重置 定时器的时间间隔。
			}

		case <-chSig: //   接收到 中断信号。停止监听事件。
			fmt.Println("recv  signal stop timer.")
			return
		}
	}
}

func (t *TimeDemo) time_struct_func() {
	fmt.Println("time.Time的一些接口和函数") // time.Time的主要接口和函数：

	/////////////////////////////////////////////////////////////////////////创建 一个time.Time 对象 有哪些方法：
	d1 := time.Date(2016, time.February /**/, 01, 23, 12, 1, 0, time.Local) //将指定年月日时分秒 转化为 time.Time
	fmt.Println("time: ", d1.String())

	fmt.Println("now().string(): ", time.Now().String(), time.Now().Local()) //获取 当前时间点的： time.Time

	//layout 有定义的变量： time.DateTime
	d2, e := time.Parse(time.DateTime, "2023-11-06 15:04:05.123") //使用time.Parse() 从指定时间字符串 按固定格式解析成 time.Time
	if e != nil {
		fmt.Println("parse fail, e: ", e)
	} else {
		fmt.Println("d2 parse time: ", d2.String())
	}

	d3, e := time.Parse("2006-01-02 15:04:05.999", "2023-11-06 15:04:05.120") // 其中小数点后的999表示，当实际中包括了0结尾，则清除0
	if e != nil {
		fmt.Println("parse fail, e: ", e)
	} else {
		fmt.Println("d3 parse time: ", d3.String())
	}

	d4, e := time.Parse("2006-01-02 15:04:05.000", "2023-11-06 15:04:05.120") // 其中小数点后的000表示，当实际中包括了0结尾，不清除 0.
	if e != nil {
		fmt.Println("d4 parse fail, e: ", e)
	} else {
		fmt.Println("d4 parse time: ", d4.String(), d4.Nanosecond())
	}

	//time.ParseInLocation()
	//time.LoadLocation("Local")
	d5, e := time.ParseInLocation(time.DateTime, "2023-11-06 15:04:05.120", /** time.UTC **/ time.Local)
	if e != nil {
		fmt.Println("d5 LoadLocation fail, e: ", e)
	} else {
		fmt.Println("d5 LoadLocation time: ", d5.String(), d5.Nanosecond())
	}

	//入参必须是时间戳，从1970开始时间段
	d6 := time.Unix(time.Now().Unix(),  int64(0)) // 将时间戳 Unix 转化为： time.Time： 从指定的秒数 转化时间点 time.Time
	if e != nil {

	} else {
		fmt.Println("time.Unix() : ",d6.String())
	}

	//入参必须是时间戳，从1970开始时间段
	d7 := time.UnixMilli(time.Now().UnixMilli()) // time.UnixMicro()
	fmt.Println("time.UnixMill(): ", d7.String()) //通过某个时间点的时间戳 获取对应的time.Time.

	////////////////////////////////////////////////另外还可以根据 其他的time.Time 和一些时间戳 运算来获取新的 time.Time
	d8 := time.Now()
	ds, _ := time.ParseDuration("4s")
	d9 := d8.Add(ds)               //////////////////////////////////// time.Time对象自己的成员方法： Add() 增加或者减少 时间 获取 新的 time.Time对象。
	fmt.Println("Add()， now: ", d8.String(), ", 4s later, time: ", d9.String())

	ds, _ = time.ParseDuration("-4s")
	d9 = d8.Add(ds)   ///////////////////////////////////////////// time.Time对象自己的成员方法： Add() 增加或者减少 时间 获取 新的 time.Time对象。
	fmt.Println("Add()， now: ", d8.String(), ", 4s before, time: ", d9.String())

	addYear, addMonth, addDay := 1, 1, 1  
	d10 := d8.AddDate(addYear, addMonth, addDay)
	fmt.Println("AddDate(), + new time.Time obj: ", d10.String())

	addYear, addMonth, addDay = -1, -1, -1 
	d11 := time.Now().AddDate(addYear, addMonth, addDay) // time.Time 对象 的成员方法： AddDate() 增加 或者减少 年， 月， 日。获取新的 time.Time 对象。
	fmt.Println("AddDate(), - new time.Time obj: ", d11.String())

	///////////////////////////////////////////////////////////////////////////////// time.Time 对象间的操作： 比较（ 在前，在后的比较）， 相减， 相加。
	t12 := time.Now()
	t13 := t12.Add(1*time.Second)
	fmt.Println("before: ", t12.Before(t13)) //时间的比较： 判断 调用方的 在参数对象的前面。
	fmt.Println("after: ", t12.After(t13)) //时间的比较： 判断 调用的time.Time对象在 参数对象time.Time 的后面。
	fmt.Println("Sub: ", t13.Sub(t12)) //时间运行，相减： 两个 time.Time对象 相减。
	fmt.Println("compare: ", t13.Compare(t12)) //时间的比较： 

	/////////////////////////////////////////////////////////////////////////////// 获取time.Time对象的年月日 数据
	ft := fmt.Sprintf("%s.000", time.DateTime)
	t14, _ := time.ParseInLocation(ft, "2006-01-02 15:04:05.812", time.Local)
	y,m,d := t14.Date()
	fmt.Printf("Date(), Year: %v, month: %v, day: %v\n", y,m,d) //返回time.Time对象的 year,month, day。
	fmt.Printf("year: %v, month: %v, day: %v, hour: %v, minute: %v, second: %v, nansecond: %v \n",
	t14.Year(), t14.Month(), t14.Day(), t14.Hour(), t14.Minute(), t14.Second(), t14.Nanosecond())

	/////////////////////////// 将time.Time对象转化 unix 的时间戳，就是比如： Unix() 时间。实际上 time.Time 和 Unix 时间存在相互转化的方法。
	fmt.Printf("Unix: %v, millisecond: %v, microsecond: %v, nansecond: %v \n", 
	time.Now().Unix(), time.Now().UnixMilli(), time.Now().UnixMicro(), time.Now().UnixNano())

	////////////////////////////////////////////////////////////////////////////// 既然存在将字符串性的时间转化为 time.Time. 也存在将time.Time 按样式格式化成字符串。
	t15 := time.Now() 
	fmt.Println("Format: ", t15.UTC().Format(time.DateTime))

}
