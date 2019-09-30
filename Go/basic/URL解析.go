package main

import (
	"fmt"
	"strings"
	"net/url"
)

// URL解析

var p = fmt.Println

func main() {
	// 解析这个 URL 示例，它包含了一个 scheme， 认证信息，主机名，端口，路径，查询参数和片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)

	p(u.User)
	p(u.User.Username())
	pd, _ := u.User.Password()
	p(pd)

	p(u.Host)
	h := strings.Split(u.Host, ":")
	p(h[0])
	p(h[1])

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])


}

/*

*/
