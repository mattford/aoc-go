package year2022

import (
	"math"
	"strconv"
	"strings"
)

type Day25 struct{}

func (p Day25) PartA(lines []string) any {
	total := 0
	for _, line := range lines {
		total += snafuToDec(line)
	}
	return decToSnafu(total)
}

func (p Day25) PartB(lines []string) any {
	return "Merry Christmas Elves!"
}

func snafuToDec(snafu string) int {
	chars := strings.Split(snafu, "")
	length := len(chars) - 1
	dec := 0
	for i, char := range chars {
		multiplier := 1
		x := length - i
		if x > 0 {
			multiplier = int(math.Pow(5, float64(x)))
		}
		var value int
		switch char {
		case "-":
			value = -1
		case "=":
			value = -2
		default:
			value, _ = strconv.Atoi(char)
		}
		dec += value * multiplier
	}
	return dec
}

func decToSnafu(dec int) string {
	snafuDigits := []rune{'0', '1', '2', '=', '-'}
	if dec == 0 {
		return ""
	}
	remainder := dec % 5

	digit := snafuDigits[remainder]

	snafu := decToSnafu((dec + 2) / 5)
	snafu += string(digit)
	return snafu
}
