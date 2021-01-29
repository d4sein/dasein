package typeguard

import (
	"reflect"
	"strconv"
	"strings"
)

// ArgumentConstructor ..
type ArgumentConstructor struct {
	To     string
	Output Output
}

// Output ..
type Output struct {
	Value string
}

// ToInt returns value of type int
func (o Output) ToInt() (output int, err error) {
	output, err = strconv.Atoi(o.Value)

	return output, err
}

// WantInt returns string for type int
func WantInt() string {
	return reflect.TypeOf(1).Name()
}

// WantArrInt returns string for type []int
func WantArrInt() string {
	return reflect.TypeOf([]int{}).Name()
}

// ToArrInt returns value of type []int
func (o Output) ToArrInt() (output []int, err error) {
	arrStrings := strings.Split(o.Value, ",")

	for i := 0; i < len(arrStrings); i++ {
		n, err := strconv.Atoi(arrStrings[i])
		if err != nil {
			return output, err
		}

		output = append(output, n)
	}

	return output, err
}
