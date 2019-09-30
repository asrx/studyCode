package main

import (
	"fmt"
	"time"
)

// 时间

var p = fmt.Println

func main() {

	// 当前时间
	now := time.Now()
	p(now)

	// 通过提供年月日等信息，你可以构建一个 time。时间总 是关联着位置信息，例如时区
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出是星期一到日的 Weekday 也是支持的
	p(then.Weekday())

	// 比较两个时间，分别测试一下是否是之前， 之后或者是同一时刻，精确到秒
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 Sub 返回一个 Duration 来表示两个时间点的间 隔时间
	diff := now.Sub(then)
	p(diff)

	// 计算出不同单位下的时间长度值
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 使用 Add 将时间后移一个时间间隔，或者使 用一个 - 来将时间前移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))

}


/*

*/
