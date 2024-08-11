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
	var m int
	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &m)
		mp := make(map[int]bool)
		rcnt := -2
		var c, cnt int
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &c)
			if rcnt == -2 {
				rcnt = -1
				if _, ok := mp[c]; !ok {
					mp[c] = false
				}
			} else if rcnt == -1 {
				if c == 0 {
					rcnt = -2
				} else {
					rcnt = c
					cnt = 0
				}
			} else {
				cnt++
				mp[c] = true
				if cnt == rcnt {
					rcnt = -2
				}
			}
		}
		for i, v := range mp {
			if !v {
				fmt.Fprintln(out, i)
			}
		}

	}
}
