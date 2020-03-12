package utils

import (
	"strings"
)

// CompareSemVers compares two semantic versions to decide the result
// Return
// 1	: 	version1 > version2
// -1	: 	version2 > version1
// 2	: 	version1 == version2
// 0	: 	error happens
//
func CompareSemVers(version1, version2 string) (int8, error) {
	if version1 == version2 {
		return 2, nil
	}
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")
	v1Length := len(v1Parts)
	v2Length := len(v2Parts)

	for v1Length < v2Length {
		v1Parts = append(v1Parts, "0")
	}

	for v2Length < v1Length {
		v2Parts = append(v2Parts, "0")
	}

	// compare version section by section
	var l = len(v1Parts)
	for i := 0; i < l; i++ {
		if v1Parts[i] == v2Parts[i] {
			continue
		}

		v1Segment, err := StrToInt8(v1Parts[i])
		if err != nil {
			return 0, err
		}
		v2Segment, err := StrToInt8(v2Parts[i])
		if err != nil {
			return 0, err
		}
		// converted result by ParseInt still in int64 format, need to coerce
		if r := int8(v1Segment) > int8(v2Segment); r == true {
			return 1, nil
		}
		return -1, nil
	}

	// should be equal if running out the loop
	return 2, nil
}
