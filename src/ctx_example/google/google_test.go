package google

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReaderCloser struct {
	strings.Reader
}

func (rc *ReaderCloser) Close() error {
	return nil
}

func NewReader(s string) *ReaderCloser {
	reader := strings.NewReader(s)

	return &ReaderCloser{*reader}
}

func TestDecodeResp(t *testing.T) {

	const jsonStream = `{
	"kind": "Ed",
	"items": [
		{
			"title": "Discover gists · GitHub",
			"link": "https://gist.github.com/"
		},
		{
			"title": "Discover gists · GitHub 1",
			"link": "https://gist.github.com/1"
		}
	]}`
	results, err := DecodeResp(NewReader(jsonStream))
	assert.Nil(t, err)
	assert.Equal(t, results, Results{
		Result{
			Title: "Discover gists · GitHub",
			URL: "https://gist.github.com/",
		},
		Result{
			Title: "Discover gists · GitHub 1",
			URL: "https://gist.github.com/1",
		},
	})
}
