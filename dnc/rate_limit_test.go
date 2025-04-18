package dnc

import (
	"slices"
	"testing"
	"time"
)

/*
TestRateLimiter tests solution(s) with the following signature and problem description:

	func IsAllowed(limitPerSecond int) bool

Given a number of allowed requests calls per second (calls/time) write an IsAllowed
function which returns false if the request should be rate limited because it exceeds the
limit and true if the request should be allowed.

For example given limit of 2 requests per second, if 3 requests are made in the first second
two calls to IsAllowed would return true and the third would return false.
*/
func TestRateLimiter(t *testing.T) {
	tests := []struct {
		limitPerSecond  int
		firstCallTimes  int
		sleep           int
		secondCallTimes int
		want            []bool
	}{
		{0, 0, 0, 0, []bool{}},
		{0, 1, 0, 0, []bool{false}},
		{0, 1, 0, 1, []bool{false, false}},
		{1, 2, 0, 2, []bool{true, false, false, false}},
		{2, 5, 0, 5, []bool{true, true, false, false, false, false, false, false, false, false}},
		{3, 5, 1, 5, []bool{true, true, true, false, false, true, true, true, false, false}},
		{10, 100, 1, 100, []bool{true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
	}

	for i, test := range tests {
		rateLimitEvents = make([]int64, 0)
		got := make([]bool, 0)
		for range test.firstCallTimes {
			got = append(got, IsAllowed(test.limitPerSecond))
		}
		time.Sleep(time.Duration(test.sleep) * time.Second)
		for range test.secondCallTimes {
			got = append(got, IsAllowed(test.limitPerSecond))
		}

		if !slices.Equal(got, test.want) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.want, got)
		}
	}
}
