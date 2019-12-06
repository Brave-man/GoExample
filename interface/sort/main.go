package main

import (
	"fmt"
	"sort"
)

// StringSlice 字符串切片
type StringSlice []string

func (p StringSlice) Len() int { return len(p) }

func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	names := []string{"D", "A", "V", "I", "D"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)

	names2 := []string{"D", "A", "V", "I", "D"}
	sort.Strings(names2)
	fmt.Println(names2)
}
