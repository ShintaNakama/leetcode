// https://jabba.cloud/20161020172918
// 動的計画法（Dynamic Programming）
package main

// 再帰を使った全探索
func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// トップダウン：メモ化再帰
var nums = make([]int, 10000)

func fib2(n int) int {
	if n <= 1 {
		return 1
	}
	if nums[n] != 0 {
		return nums[n]
	}

	r := fib2(n-1) + fib2(n-2)
	nums[n] = r
	return r
}

// ボトムアップ：漸化式ループ
func fib3(n int) int {
	if n == 0 {
		return 1
	}

	var (
		fib1 = 1
		fib2 = 1
		fib3 = 1
	)

	for i := 1; i < n; i++ {
		fib3 = fib1 + fib2
		fib1 = fib2
		fib2 = fib3
	}
	return fib2
}
