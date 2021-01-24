package typeguard

import (
	"strconv"
)

// ToInteger ..
func ToInteger(input string) (output int, err error) {
	output, err = strconv.Atoi(input)

	return output, err
}
