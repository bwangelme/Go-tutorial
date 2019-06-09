package testify_exam

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// `require.True` 失败后会将测试 TestEx1 停掉，不会再执行后面的 require.Equal 断言
// `assert.True` 失败后并不会将 TestEx2 停掉，在后面的断言中，Add 依然会被执行，
// 所以 Add 会在 TestEx2 中执行一次， count 的最终结果是1

var count int32 = 0

func Add(a, b int) int {
	atomic.AddInt32(&count, 1)
	return a + b
}

func TestEx1(t *testing.T) {
	require.True(t, false)
	require.Equal(t, Add(1, 1), 2)
}

func TestEx2(t *testing.T) {
	assert.True(t, false)
	assert.Equal(t, Add(1, 1), 2)
	assert.Equal(t, count, int32(1))
}

// 下面是程序的运行结果
//>>> go test -v testify_exam/require_test.go                                                                                                                      23:30:19 (06-02)
//=== RUN   TestEx1
//--- FAIL: TestEx1 (0.00s)
//Error Trace:    require_test.go:19
//Error:          Should be true
//=== RUN   TestEx2
//--- FAIL: TestEx2 (0.00s)
//Error Trace:    require_test.go:24
//Error:          Should be true
//FAIL
//FAIL    command-line-arguments  0.018s
