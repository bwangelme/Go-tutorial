package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	cmd = exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("Some input")
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", stdOut.String())
}
