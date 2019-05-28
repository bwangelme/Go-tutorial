package b64_example

import (
	"encoding/base64"
	"fmt"
	"github.com/d5/tengo/assert"
	"testing"
)

func TestNoPadding(t *testing.T) {
	data := []byte{'A', 0}

	encoded := base64.RawStdEncoding.EncodeToString(data)
	decodedData, err := base64.RawStdEncoding.DecodeString(encoded)
	assert.Nil(t, err)

	fmt.Println(encoded)
	fmt.Println(decodedData)
}
