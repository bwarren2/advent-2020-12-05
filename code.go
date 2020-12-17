package advent202001205

import (
	"bufio"
	"os"
	"sort"
)

// Seat represents a flight seat
type Seat struct {
	LoRow, HiRow, LoCol, HiCol int64
}

// NewSeat initializes a Seat from a binary space partition string
func NewSeat(bsp string) *Seat {
	seat := new(Seat)
	seat.HiRow = 127
	seat.HiCol = 7

	for _, char := range bsp {

		if char == 'F' {
			diff := seat.HiRow - seat.LoRow
			seat.HiRow = seat.LoRow + (diff-diff%2)/2
		} else if char == 'B' {
			diff := seat.HiRow - seat.LoRow
			seat.LoRow = seat.LoRow + (diff+diff%2)/2
		} else if char == 'L' {
			diff := seat.HiCol - seat.LoCol
			seat.HiCol = seat.LoCol + (diff-diff%2)/2
		} else if char == 'R' {
			diff := seat.HiCol - seat.LoCol
			seat.LoCol = seat.LoCol + (diff+diff%2)/2
		}

	}

	return seat
}

// Col gets the column for the seat
func (s *Seat) Col() int64 {
	if s.LoCol == s.HiCol {
		return s.LoCol
	}
	panic("Couldn't get a col")
}

// Row gets the row for the seat
func (s *Seat) Row() int64 {
	if s.LoRow == s.HiRow {
		return s.LoRow
	}
	panic("Couldn't get a row")
}

// ID gets the Seat ID
func (s *Seat) ID() int64 {
	return 8*s.Row() + s.Col()
}

// LinesFromFile gets an array of cases
func LinesFromFile(filename string) (result []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("What??")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

// Part1 answers part 1
func Part1(filename string) int64 {
	var max int64
	for _, str := range LinesFromFile(filename) {
		seat := NewSeat(str)
		if seat.ID() > max {
			max = seat.ID()
		}
	}
	return max
}

// Part2 answers part 2
func Part2(filename string) int64 {
	var nums []int64
	for _, str := range LinesFromFile(filename) {
		seat := NewSeat(str)
		nums = append(nums, seat.ID())
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	// fmt.Println(nums)
	offset := nums[0]
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx]-offset != int64(idx) {
			return int64(idx) + offset
		}
	}
	return -1
}
