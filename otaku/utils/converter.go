package utils

import (
	"strconv"
)

// StrToInt8 converts from string to type int8
func StrToInt8(s string) (int8, error) {
	num, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return -1, err
	}

	// converted result by ParseInt still in int64 format, need to coerce
	numIn8 := int8(num)

	return numIn8, nil
}
