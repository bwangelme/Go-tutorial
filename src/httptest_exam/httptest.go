package httptest_exam

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func Greet() string {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(greeting)
}
