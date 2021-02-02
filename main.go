package main

import (
	"fmt"
	"bufio"
	"os"
  "net/url"
)

func main() {
  fmt.Println("===Input Percent Encoded URL!===")
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  input := stdin.Text()
  out, _ := url.QueryUnescape(input)
  fmt.Println("")
  fmt.Println("===Here you are!===")
  fmt.Println(out)
}
