package testify_exam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	// 测试相等
	assert.Equal(t, 123, 123, "they should be equal")

	// 测试不相等
	assert.NotEqual(t, 123, 456, "they should not be equal")

	var obj chan struct{}
	// 测试 nil，常用于错误值的判断
	assert.Nil(t, obj)

	obj2 := struct {
		Value string
	}{
		Value: "Something",
	}
	// 测试不为 nil，常用于你希望返回的结果值拥有某些东西的时候
	if assert.NotNil(t, obj2) {
		// 现在我们知道 obj2 的值不是 nil了，我们可以在这个基础上**安全地**访问它的字段。
		assert.Equal(t, "Something", obj2.Value)
	}
}

func TestEasyAssert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(123, 123, "they should be equal")

	assert.NotEqual(123, 456, "they should not be equal")

	var obj chan struct{}
	assert.Nil(obj)

	obj2 := struct {
		Value string
	}{
		Value: "Something",
	}
	if assert.NotNil(obj2) {
		assert.Equal("Something", obj2.Value)
	}
}
