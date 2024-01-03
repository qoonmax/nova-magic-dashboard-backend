package utils

import (
	"strconv"
	"strings"
)

func ValidateVersion(minVersion string, currentVersion string) bool {
	minVersionArr := strings.Split(minVersion, ".")
	currentVersionArr := strings.Split(currentVersion, ".")

	for i := 0; i < len(minVersionArr) || i < len(currentVersionArr); i++ {
		var minNum, currentNum int
		var err error

		if len(minVersionArr) >= i {
			minNum, err = strconv.Atoi(minVersionArr[i])
			if err != nil {
				panic(err)
			}
		} else {
			minNum = 0
		}

		if len(currentVersionArr) >= i {
			currentNum, err = strconv.Atoi(currentVersionArr[i])
			if err != nil {
				panic(err)
			}
		} else {
			currentNum = 0
		}

		if currentNum > minNum {
			break
		} else if currentNum == minNum {
			continue
		} else {
			return false
		}
	}
	return true
}
