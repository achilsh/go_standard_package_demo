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
}

func (t *TimeDemo) after() { //从现在 多少时间之后的点。
	fmt.Println("call time.After(), ", time.Now().Unix())
	//等待一段时间后向 返回的chan发送一个通知； 调用方 阻塞等待chan通知。可以用来等待将来的某个时间点。
	c := time.After(2*time.Second) //下一个 n time的时间点。
	select {
	case <-c:
		fmt.Println("等待一段时间后，当前时间已经到了。") //可用于超时计时器。
	}
	fmt.Println("end time: ", time.Now().Unix())

	// sleep:
	fmt.Println("time.Sleep(1*time.Second)")
	time.Sleep(1*time.Second)

}


func( t* TimeDemo) tick() {
	//主要包装了 time.NewTicker , 
	fmt.Println("run time.Tick()")
	tt := time.Tick(2*time.Second) //返回一个chan, 系统会定时往chan 发送一个信号。 调用方 可以不断的消费chan 实现定时器。

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
func (t *TimeDemo)parseDuration() {
	fmt.Println("run time.ParseDuration()")
	h, _ := time.ParseDuration("1h") // parses a duration string.  主要用来将字符串的时间 转化 time.Duration. 其中：字符串中出现的单位有： "ns", "us" (or "µs"), "ms", "s", "m", "h"
	fmt.Printf("hour: %#v\n", h/time.Hour)
	s, _ := time.ParseDuration("3s")
	fmt.Printf("second: %#v\n", s/time.Second)


}

func (t *TimeDemo) since() {
	fmt.Println("time.Since(t)") //主要用来计算之前的某个时间点 到现在时刻的 时间 time.Duration .通常是用于计算时间差。 参考点是过去和当前time.Now().
	before := time.Now() 
	<- time.After(2*time.Second)
	elapsed := time.Since(before) //计算过去到现在的时间差。
	fmt.Printf("%v\n", elapsed)
}

func (t *TimeDemo) hour() {
	fmt.Println("time.Duration.Hour()") //返回指定时间的一共多少hour。 
	d, _ := time.ParseDuration("2h")
	
	fmt.Println("hour: ", d.Hours()) //time.Duration 一共折算成多少 hour
	fmt.Println("minute: ", d.Minutes()) //time.Duration 一共 折算成多少 分钟
	fmt.Println("second: ", d.Seconds()) //time.Duration 一共折算成多少 秒
	fmt.Println("millSecond: ", d.Milliseconds())//time.Duration 一共折算成多少 毫秒

	d2, _ := time.ParseDuration("1h15m30.918273645s")
	fmt.Println(d2.Hours(), d2.Minutes(), d2.Seconds(), d2.String())
}

func (t *TimeDemo)string_call() {
	fmt.Println("time.Duration.String()")  //将 time.Duration 的值转化为 字符串。比如： 2*time.Hour 可以转化 2h ，他的反操作是time.ParseDuration()
	a := time.Hour *1 + time.Minute *2 + time.Second * 3 + time.Millisecond * 4 + time.Microsecond * 5
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

func (t* TimeDemo) ticker() {
	fmt.Println("time.Ticker") //定时器，  他内部保存了一个chan, 系统会定期往 chan 发送 tick. 调用方可 不断的从 chan 读取tick 实现定时。
	tk := time.NewTicker(1*time.Second) //创建一个指定定时间隔的 ticker.
	defer tk.Stop() //停掉 该定时器。

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
		
	i  := 0
	for {
	select {
		case <-tk.C:  			//此时接收到定时器 通知。 因为系统往定时器的chan 发送了tick。
			fmt.Println("tick arrive, you can press ctrl +c stop.")
			i++
			if i == 5 {
				fmt.Println("call Reset(d): ")
				tk.Reset(2*time.Second) // 重置 定时器的时间间隔。
			}
			
		case <-chSig: //   接收到 中断信号。停止监听事件。
			fmt.Println("recv  signal stop timer.")
			return 
	}	
	}
}