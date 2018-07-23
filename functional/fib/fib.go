package fib

func Fibonacci() func() int{
	first := 0
	second := 1
	return func() int {
		cur := first + second
		first = second
		second = cur
		return cur
	}
}
