package utils

import (
	"strconv"
)

/*
StringToUint is a utility function that converts a string to an unsigned 32-bit integer. It parses
the input string and returns the corresponding uint32 value. If parsing fails, it returns 0. It is
used for converting route parameters or other string representations of unsigned integers.
*/
func StringToUint(s string) uint32 {
	i, _ := strconv.Atoi(s)
	return uint32(i)
}
