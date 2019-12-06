// 通过类型断言来判断w的动态类型是否满足一个新接口
package main

import (
	"io"
	"os"
)

func writerString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriterString(string) (n int, err error)
	}

	if sw, ok := w.(stringWriter); ok {
		return sw.WriterString(s)
	}
	return w.Write([]byte(s))
}

func writerHeader(w io.Writer, contentType string) error {
	if _, err := writerString(w, "Content-Type: "); err != nil {
		return err
	}

	if _, err := writerString(w, contentType); err != nil {
		return err
	}

	return nil
}

func main() {
	w := os.Stdout
	_ = writerHeader(w, "Application/json")
}
