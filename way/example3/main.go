// nil是一个合法的接受者
package main

import (
	"fmt"

	"goCode/review/way/example3/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "A")
	m.Add("item", "B")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil // nil是一个合法的接受者
	fmt.Println(m.Get("item"))
	m.Add("h", "1") // 宕机: 赋值给一个空map类型
}
