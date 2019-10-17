package memo_test

import (
	"memo"
	"memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
