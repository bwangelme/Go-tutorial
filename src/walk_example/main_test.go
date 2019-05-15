package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	// path 是路径，f.Name 是文件名或者目录名
	fmt.Printf("Visited: %s, %s\n", path, f.Name())
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "."
	}

	err := os.Chdir(root)
	if err != nil {
		log.Fatalln(err)
	}

	err = filepath.Walk(".", visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	// Output
	// Visited: ., .
	// Visited: dir2, dir2
	// Visited: dir2/dir3, dir3
	// Visited: dir2/file1, file1
	// Visited: dir2/file2, file2
	// Visited: file1, file1
	// filepath.Walk() returned <nil>

}
