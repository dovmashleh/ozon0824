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
	var n, p int
	var price int
	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n, &p)
		sum := 0
		for sold := 0; sold < n; sold++ {
			fmt.Fscan(in, &price)
			lost := price*p - (price*p/100)*100
			sum += lost
		}
		fmt.Fprintf(out, "%.2f\n", float32(sum)/100)
	}
}
