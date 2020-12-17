package advent202001205_test

import (
	"fmt"
	"testing"

	advent "github.com/bwarren2/advent20201205"
)

func TestNewSeat(t *testing.T) {
	tcs := []struct {
		name, input string
		want        int64
	}{
		{name: "Example 1", input: "BFFFBBFRRR", want: 567},
		{name: "Example 2", input: "FFFBBBFRRR", want: 119},
		{name: "Example 3", input: "BBFFBBFRLL", want: 820},
	}
	for _, tc := range tcs {
		seat := advent.NewSeat(tc.input)
		got := seat.ID()
		if got != tc.want {
			t.Errorf("Got the wrong answer %v in %v", got, tc)
		}
	}
}

func TestPart1(t *testing.T) {
	fmt.Println(advent.Part1("input.txt"))
}
func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("input.txt"))
}
