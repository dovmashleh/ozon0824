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
	var n int

	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n)
		l := make([]int, n)
		r := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &(l[i]))
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &(r[i]))
		}
		prod := 1
		for i := 0; i < n; i++ {
			c := r[i]/(i+1) - l[i]/(i+1)
			if l[i]%(i+1) == 0 {
				c++
			}
			prod *= c
			prod %= 1000000007
		}
		fmt.Fprintln(out, prod)
	}
}
