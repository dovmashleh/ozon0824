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

	var n, m int
	leetEx := []string{}
	leetWant := []string{}
	existing := make(map[int][]string)
	fmt.Fscan(in, &n)
	var inp string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &inp)
		sum := computeSum(inp)
		existing[sum] = append(existing[sum], inp)
		leetEx = append(leetEx, inp)
	}
	fmt.Fscan(in, &m)
	results := make([]bool, 0, m)
NEWBIELOOP:
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &inp)
		leetWant = append(leetWant, inp)
		sum := computeSum(inp)
		candidates, ok := existing[sum]
		if ok {
			for _, c := range candidates {
				if areSimilar(c, inp) {
					fmt.Fprintln(out, 1)
					results = append(results, true)
					continue NEWBIELOOP
				}
			}
		}
		fmt.Fprintln(out, 0)
		results = append(results, false)
	}
	fmt.Println("input: n=", n, "m=", m, "existing=", leetEx, "wanted=", leetWant)
	fmt.Println("output: ", results)
}
func computeSum(s string) int {
	var sum int
	for _, c := range s {
		sum += int(c)
	}
	return sum*1000 + len(s)
}

func areSimilar(a string, b string) bool {
	/*if len(a) != len(b) {
		return 0
	}*/
	foundPerm := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if foundPerm {
				return false
			}
			if i < len(a)-1 && a[i+1] == b[i] && a[i] == b[i+1] {
				foundPerm = true
				i++
			} else {
				return false
			}
		}
	}
	return true
}
