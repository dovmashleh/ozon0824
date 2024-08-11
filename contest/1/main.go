package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tests int
	var n string

	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n)
		asbyte := []byte(n)
		pos := -1
		prev := byte(9)

		for j, v := range asbyte {
			if v-48 > prev {
				pos = j - 1
				break
			}
			prev = v - 48
		}
		if pos == -1 {
			pos = len(asbyte) - 1
		}

		asbyte = append(asbyte[:pos], asbyte[pos+1:]...)
		if len(asbyte) == 0 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, string(asbyte))
		}
	}
}
