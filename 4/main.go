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
	var n, v int
	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n)
		mx := 0
		a := 0
		b := 0
		l := 0
		s := 0
		fmt.Fscan(in, &a)
		changed := false
		prev := a
		for r := 1; r < n; r++ {
			fmt.Fscan(in, &v)
			if a != v {
				if b == 0 {
					b = v
					l = r
					prev = b
					continue
				}
			}

			if v != a && v != b {
				a = b
				b = v
				tmx := r - s
				s = l
				if tmx > mx {
					mx = tmx
				}
				changed = true
				l = r
				prev = v
				continue
			}
			if prev != v {
				l = r
			}
			if v != b && v == a && b != 0 {
				a, b = b, a
			}
			prev = v
		}

		if changed {
			tmx := n - s
			if tmx > mx {
				mx = tmx
			}
			fmt.Fprintln(out, mx)
		} else {
			fmt.Fprintln(out, n)
		}
	}
}
