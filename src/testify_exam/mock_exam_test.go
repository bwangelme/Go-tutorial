package testify_exam

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) DoRequest(url string) (*http.Response, error) {
	_ = m.Called()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World")
		return
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TestMockSomething(t *testing.T) {
	testObj := new(MyMockedObject)

	testObj.On("DoSomething", "http://www.baidu.com").Return()
	err := UpdateFile("/tmp/hello.txt", testObj)
	assert.NoError(t, err)
}
