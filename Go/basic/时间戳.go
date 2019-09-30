package main

import (
	"fmt"
	"time"
)

// 时间戳

var p = fmt.Println

func main() {

	// 使用带 Unix 或者 UnixNano 的 time.Now 来获取从自协调世界时 起到现在的秒数或者纳秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	p(now)


	// 注意 UnixMillis 是不存在的，
	// 所以要得到毫秒数的话， 
	// 你要自己手动的从纳秒转化一下
	millis := nanos / 1000000
	p(secs)
	p(millis)
	p(nanos)

	// 可以将协调世界时起的整数秒或者纳秒转化到相应的时间
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))


}


/*

*/
