package main

import "fmt"

//https://qiita.com/drken/items/23a4f604fa3f505dd5ad
// この記事が参考になる特にhttps://qiita.com/drken/items/23a4f604fa3f505dd5ad#n--5-%E3%81%AE%E3%81%A8%E3%81%8D
func addRecurcive(n int) int {
	fmt.Println(n)
	if n == 0 {
		return 0
	}

	return n + addRecurcive(n-1)
}
