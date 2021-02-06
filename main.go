package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe != 0 {
		readByPipe()
	} else {
		readByInteractive()
	}
}

func readByPipe() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		decoded, _ := url.QueryUnescape(string(line))
		fmt.Println(decoded)
	}
}

func readByInteractive() {
	fmt.Println("===Input Percent Encoded URL!===")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := stdin.Text()
	out, _ := url.QueryUnescape(input)
	fmt.Println("")
	fmt.Println("===Here you are!===")
	fmt.Println(out)
}
