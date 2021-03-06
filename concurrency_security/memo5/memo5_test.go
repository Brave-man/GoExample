package memo5_test

import (
	"testing"

	"goCode/review/concurrency_security/memo5"
	"goCode/review/concurrency_security/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Concurrent(t, m)
}
