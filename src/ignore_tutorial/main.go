package main

import (
	"fmt"
	"log"

	ignore "github.com/sabhiram/go-gitignore"
)

const (
	IsDir int = iota + 1
	IsNotExist
)

func main() {
	fmt.Println(IsDir, IsNotExist)
	lines := []string{"abc/def", "venv/", "b"}
	giPattern, err := ignore.CompileIgnoreLines(lines...)
	if err != nil {
		log.Fatal(err)
	}
	//assert.Nil(test, error, "error from CompileIgnoreLines should be nil")

	res := giPattern.MatchesPath("venv/")
	fmt.Println(res)
}
