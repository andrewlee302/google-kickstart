package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ti := 1; ti <= t; ti++ {
		var l int
		var s string
		fmt.Scanf("%d", &l)
		fmt.Scanf("%s", &s)
		subarrayLen := (l + 1) / 2
		max := []int{0, -1}
		windowSum := 0
		for i := 0; i < l; i++ {
			if i < subarrayLen {
				windowSum += int(s[i] - '0')
			} else {
				windowSum += int(s[i] - '0')
				windowSum -= int(s[i-subarrayLen] - '0')
			}
			if windowSum > max[0] {
				max = []int{windowSum, i}
			}
		}
		fmt.Printf("Case #%d: %d\n", ti, max[0])
	}
}
