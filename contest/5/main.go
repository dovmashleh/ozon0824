package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	isfirst := true
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n)
		_, _ = in.ReadString('\n')
		stack := make([]string, 0)
		printed := false
		for i := 0; i < n; i++ {
			outp, _ := in.ReadString('\n')
			arr := strings.Fields(outp)
			depth := strings.Index(outp, arr[0]) / 4
			key := arr[0][:len(arr[0])-1]
			if depth < len(stack) {
				stack = stack[:depth]
				printed = false
			}
			if len(arr) == 1 {
				stack = append(stack, key)
				printed = false
			} else {
				if !printed {
					if !isfirst {
						fmt.Fprintln(out)
					} else {
						isfirst = false
					}
					if len(stack) > 0 {
						fmt.Fprintf(out, "[%s]\n", strings.Join(stack, "."))
					}
				}
				fmt.Fprintf(out, "%s = %s\n", key, arr[1])
				printed = true
			}
		}

	}
}
