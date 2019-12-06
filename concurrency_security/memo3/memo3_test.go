package memo3_test

import (
	"testing"

	"goCode/review/concurrency_security/memo3"
	"goCode/review/concurrency_security/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Concurrent(t, m)
}
