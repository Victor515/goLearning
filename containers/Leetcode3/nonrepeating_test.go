package main

import "testing"

/*
test-related command
go test: run all tests in current directory
go test -bench . : run all benchmark tests in current directory
go test -bench . -cpuprofile cpu.out: output cpu data for becnhmark tests
go test pprof: analyze cpu data(this will open a interative terminal)
After running go test pprof: web -- visualize cpu data
go test -cover: check coverage of tests
go test --coverageprofile=c.out: output coverage profile

 */

func TestSubstr(t *testing.T){
	// testcases
	tests := []struct {
		s string
		ans int
	}{
		// Normal Cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge Cases
		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"abcabcabcd", 4},

		// International Support
		{"吃葡萄不吐葡萄皮", 5},
		{"这里是慕课网", 6},
		{"一二三三二一", 3},
		{"黑化肥发灰会挥发灰化肥挥发会发黑", 7},
	}

	// run test
	for _, tt := range tests{
		actual := solveInt(tt.s)
		if actual != tt.ans{
			t.Errorf("Got %d for input: %s; expected: %d\n", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s := "黑化肥发灰会挥发灰化肥挥发会发黑"
	for i := 0; i < 13; i++{
		s = s + s
	}
	ans := 7

	b.Logf("len(s) = %d\n", len(s))

	// leave out data preparation time
	b.ResetTimer()

	for i := 0; i < b.N; i++{
		actual := solveInt(s)
		if actual != ans{
			b.Errorf("Got %d for input: %s; expected: %d\n", actual, s, ans)
		}
	}
}
