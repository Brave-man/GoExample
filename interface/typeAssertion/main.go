// 类型断言是作用在接口值上的操作，类似于x.(T)
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	// 1.如果T是一个具体类型，类型断言会去检查x的动态类型是否是T？
	// 如果检查成功，类型断言的结果就是x的动态值，类型就是T。
	// 可以理解为：类型断言就是从x中把具体类型T的值提取出来的操作。
	f := w.(*os.File)
	fmt.Printf("%T\n", f)
	if c, ok := w.(*bytes.Buffer); ok {
		fmt.Println(c)
	} else {
		fmt.Println("check fail")
	}

	// 2.如果T是一个接口类型，那么类型断言检查x的动态值是否满足T？
	// 如果检查成功，动态值并没有提取出来，结果仍然是一个接口值，接口值的类型和值部分也没有发生变更，
	// 只是结果的类型为接口类型T.
	// 可以理解为：类型断言是一个接口值表达式，从一个接口类型变为拥有另外一套方法的接口类型（通常方法数量增多），
	// 但是保留了接口值中的动态类型和动态值部分。
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", rw)
	// 此时rw接口变为ReadWriter, 拥有实现ReadWriter接口中的所有方法

	// 3.无论类型T作为具体类型还是接口类型，当操作数x是一个空接口值，类型断言都要失败
}
