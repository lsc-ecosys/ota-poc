package utils

import (
	"errors"
	"strconv"
	"strings"
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

// ConvertVersionToInt8Array split semVer string and return int8 array
func ConvertVersionToInt8Array(version string) ([]int8, error) {
	vStrArr := strings.Split(version, ".")
	var versionInt8Array = []int8{}
	if len(vStrArr) == 0 {
		return []int8{}, errors.New("empty version")
	}
	for i, s := range vStrArr {
		var err error
		versionInt8Array[i], err = StrToInt8(s)
		if err != nil {
			return []int8{}, errors.New("conversion error")
		}
	}

	return versionInt8Array, nil
}

// ConvertIntSemVerToString convert each section of semVer from int8 to string
// And return concatenated version string
func ConvertIntSemVerToString(mj, mn, pc int8) string {
	verStr := strconv.Itoa(int(mj)) + "." + strconv.Itoa(int(mn)) + "." + strconv.Itoa(int(pc))
	return verStr
}
