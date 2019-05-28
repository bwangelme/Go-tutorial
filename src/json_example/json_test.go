package json_example

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func TestMarshal(t *testing.T) {

	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)

	assert.Nil(t, err)
	assert.Equal(t, b, []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`))
}

func TestMarshalMap(t *testing.T) {
	// 结构体只会编码 Exported 字段，map会编码所有字段
	v := map[string]interface{}{
		"xff":    "whf",
		"age":    22,
		"gender": "female",
	}

	b, err := json.Marshal(v)
	assert.Nil(t, err)
	assert.Equal(t, b, []byte(`{"age":22,"gender":"female","xff":"whf"}`))
}

func TestUnmarshal(t *testing.T) {
	data := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m Message
	var expectedM = Message{
		Name: "Alice",
		Body: "Hello",
		Time: 1294706395881547000,
	}

	err := json.Unmarshal(data, &m)
	assert.Nil(t, err)
	assert.Equal(t, m, expectedM)
}

func TestUnmarshalOrder(t *testing.T) {
	data := []byte(`{"Name1":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m struct {
		Name  string
		NAME1 string
	}

	err := json.Unmarshal(data, &m)
	assert.Nil(t, err)
	assert.Equal(t, m.NAME1, "Alice")
	assert.Equal(t, m.Name, "")

	var m1 struct {
		Name  string `json:"Name1"`
		NAME1 string
	}

	err = json.Unmarshal(data, &m1)
	assert.Nil(t, err)
	assert.Equal(t, m1.Name, "Alice")
	assert.Equal(t, m1.NAME1, "")
}

func TestUnmarshalArbitraryData(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia", 2]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	assert.Nil(t, err)

	m, ok := f.(map[string]interface{})
	assert.True(t, ok)

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				switch vv := u.(type) {
				case string:
					fmt.Println(i, "is string", vv)
				case float64:
					fmt.Println(i, "is float64", vv)
				}
			}
		default:
			fmt.Println(k, "is of type I don't know how to handle")
		}
	}
}

func TestUnmarshalReferenceTypes(t *testing.T) {
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}

	var m FamilyMember

	data := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	err := json.Unmarshal(data, &m)

	assert.Nil(t, err)
	assert.Equal(t, m.Name, "Wednesday")
	assert.Equal(t, m.Age, 6)
	assert.Equal(t, m.Parents, []string{"Gomez", "Morticia"})
}

// 下面的测试函数测试了 json.Encoder 的 SetEscapeHTML 函数
func TestMarshalEscapeHTML(t *testing.T) {
	var val = map[string]string{
		"string": "<p>文本&文本</p>",
	}

	var data = bytes.NewBuffer([]byte{})
	var enc = json.NewEncoder(data)

	err := enc.Encode(val)
	assert.Nil(t, err)
	// 注意 JSON 输出最后有个换行，所以 Expected 值中也应该加上换行
	assert.Equal(t, data.Bytes(), []byte(`{"string":"\u003cp\u003e文本\u0026文本\u003c/p\u003e"}
`))

	data.Reset()
	enc.SetEscapeHTML(false)
	err = enc.Encode(val)
	assert.Nil(t, err)

	assert.Equal(t, data.Bytes(), []byte(`{"string":"<p>文本&文本</p>"}
`))
}

func TestMarshalByteArr(t *testing.T) {
	content := []byte("窗前莺并语，窗外燕双飞")
	var val = map[string]interface{}{
		"data": content,
		"b":    'A',
	}

	b64Content := base64.StdEncoding.EncodeToString(content)

	data, err := json.Marshal(val)
	assert.Nil(t, err)
	assert.Equal(t, string(data), fmt.Sprintf(`{"b":65,"data":"%s"}`, b64Content))
}

func TestMarshalName(t *testing.T) {
	var val = struct {
		ContentType string
	}{
		ContentType: "text/html",
	}

	data, err := json.Marshal(val)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `{"ContentType":"text/html"}`)
}

func TestMarshalStructFieldTag(t *testing.T) {
	// 注意，首先要让字段名称为导出名称（即以大写字母开头，struct field tag 才会生效）
	var val = struct {
		// 指定 JSON 对象中的 key 为 myName1
		F1 int `json:"myName1"`

		// 指定 JSON 对象中的 key 为 myName2，且为空值时会被忽略
		F2 int `json:"myName2,omitempty"`

		// 未指定 JSON 对象的 key 名称，所以它是 F3,指定编码选项，为空值时将会被忽略
		F3 int `json:",omitempty"`

		// 这个字段将会被忽略
		F4 int `json:"-"`

		// 将字段名指定成 -
		F5 int `json:"-,"`
	}{
		F1: 0,
		F2: 0,
		F3: 3,
		F4: 4,
		F5: 5,
	}

	data, err := json.Marshal(val)
	assert.Nil(t, err)

	assert.Equal(t, string(data), `{"myName1":0,"F3":3,"-":5}`)
}

func TestMarshalStringTag(t *testing.T) {
	// 注意 struct field tag 定义出错的时候(例如字段名称重复)，Marshal 的 err 返回的仍然是空值，但data为`{}`
	var val = struct {
		B1 bool `json:"boolVal1,string"`
		B2 bool `json:"boolVal2"`
	}{
		B1: false,
		B2: false,
	}

	data, err := json.Marshal(val)
	assert.Nil(t, err)

	assert.Equal(t, string(data), `{"boolVal1":"false","boolVal2":false}`)
}

func TestMarshalAnonymousStruct(t *testing.T) {
	type Engine struct {
		Power int
		Code  int
	}

	type Tires struct {
		Number int
	}

	type Bar interface{}

	var val = struct {
		Engine
		Tires `json:"Tires"`
		Bar
	}{
		Engine{1, 2},
		Tires{2},
		233,
	}

	data, err := json.Marshal(val)
	assert.Nil(t, err)

	assert.Equal(t, string(data), `{"Power":1,"Code":2,"Tires":{"Number":2},"Bar":233}`)
}

func TestMarshalCyclicStruct(t *testing.T) {
	type Engine struct {
		Power int
		Code  int
		Tires struct {
			Number int
			Engine
		}
	}

	var val = new(Engine)

	data, err := json.Marshal(val)
	assert.Nil(t, err)
	assert.Equal(t, data, ``)
	// >>> go test json_example                                                                                                                                         00:02:01 (05-28)
	// # json_example [json_example.test]
	//panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x70 pc=0x17e3d29]
	//
	// goroutine 1 [running]:
	// cmd/compile/internal/gc.dowidth(0xc0004cf920)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:175 +0xa9
	// cmd/compile/internal/gc.widstruct(0xc0004cf860, 0xc0004cf860, 0x0, 0x1, 0x3)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:95 +0xc6
	// cmd/compile/internal/gc.dowidth(0xc0004cf860)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:340 +0x5e9
	// cmd/compile/internal/gc.widstruct(0xc0004cf920, 0xc0004cf920, 0x0, 0x1, 0xc000448d01)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:95 +0xc6
	// cmd/compile/internal/gc.dowidth(0xc0004cf920)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:340 +0x5e9
	// cmd/compile/internal/gc.widstruct(0xc0004cf8c0, 0xc0004cf8c0, 0x0, 0x1, 0xc0004cf8c0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:95 +0xc6
	// cmd/compile/internal/gc.dowidth(0xc0004cf8c0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:340 +0x5e9
	// cmd/compile/internal/gc.resumecheckwidth()
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/align.go:450 +0x51
	// cmd/compile/internal/gc.typecheckdef(0xc00042fe00)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:3927 +0x8e3
	// cmd/compile/internal/gc.typecheck1(0xc00042fe00, 0x4, 0x0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:376 +0xc387
	// cmd/compile/internal/gc.typecheck(0xc00042fe00, 0x4, 0x0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:299 +0x6f2
	// cmd/compile/internal/gc.typecheck1(0xc000437480, 0x1, 0x0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:2204 +0x5446
	// cmd/compile/internal/gc.typecheck(0xc000437480, 0x1, 0x0)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:299 +0x6f2
	// cmd/compile/internal/gc.typecheckslice(0xc000424480, 0x6, 0x8, 0x1)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/typecheck.go:117 +0x50
	// cmd/compile/internal/gc.Main(0x1a73428)
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/internal/gc/main.go:545 +0x29a5
	// main.main()
	// /usr/local/Cellar/go/1.12.4/libexec/src/cmd/compile/main.go:51 +0xad
	// FAIL    json_example [build failed]
}
