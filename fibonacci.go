package main

import "fmt"

func fibonacci() func(int) int {
	sum := 0
	j := 0
	k := 0
	return func(x int) int {
		if x == 0 {
			return sum
		}
		if x == 1 {
			j = 1
			sum = j + k
			return sum
		}
		sum = j + k
		k = j
		j = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
