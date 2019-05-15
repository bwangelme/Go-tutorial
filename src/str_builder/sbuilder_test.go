package str_builder

import (
	"bytes"
	"strings"
	"testing"
)

// bytes.Buffer 将 buf 转换成字符串类型
// strings.Builder 没有复制数据，直接将 buf 地址存储的数据转换成了 string 类型 `*(*string)(unsafe.Pointer(&b.buf))`
// 所以构造字符串 strings.Builder 要比 bytes.Buffer 要快很多

func BenchmarkBuilder(b *testing.B) {
	var str strings.Builder

	str.WriteString("xff")

	for i := 0; i < b.N; i++ {
		str.String()
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer

	buf.WriteString("xff")

	for i := 0; i < b.N; i++ {
		buf.String()
	}
}

//>>> go test -bench=. -run='^$' tutorial/str_builder                                                                                                                                                                          23:32:03 (04-24)
//goos: darwin
//goarch: amd64
//pkg: tutorial/str_builder
//BenchmarkBuilder-4      2000000000               0.38 ns/op
//BenchmarkBuffer-4       200000000                7.85 ns/op
//PASS
//ok      tutorial/str_builder    3.164s
