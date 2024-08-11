package main

func checkSimilar(n, m int, exist []string, want []string) []bool {
	existing := make(map[int][]string)
	for i := 0; i < n; i++ {
		inp := exist[i]
		sum := computeSum(inp)
		existing[sum] = append(existing[sum], inp)
	}
	results := make([]bool, 0, m)
NEWBIELOOP:
	for i := 0; i < m; i++ {
		inp := want[i]
		sum := computeSum(inp)
		candidates, ok := existing[sum]
		if ok {
			for _, c := range candidates {
				if areSimilar(c, inp) {
					results = append(results, true)
					continue NEWBIELOOP
				}
			}
		}
		results = append(results, false)
	}
	return results
}
func лcomputeSum(s string) int {
	var sum int
	for _, c := range s {
		sum += int(c)
	}
	return sum*1000 + len(s)
}

func лareSimilar(a string, b string) bool {
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
