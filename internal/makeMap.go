package internal

import (
	"regexp"
	"strings"
)

func MakeMap(d *Ascii) {
	tempChar := make(map[rune][]string)
	tempArr := []string{}
	firstSymbol := ' '

	symbol := regexp.MustCompile(`\r`)
	strText := symbol.ReplaceAllString(string(d.DataFromBannerFile), "")
	sliceOfText := strings.Split(string(strText), "\n")

	for _, val := range sliceOfText {
		if len(val) != 0 {
			tempArr = append(tempArr, val)
		} else if len(val) == 0 && len(tempArr) != 0 {
			tempChar[firstSymbol] = tempArr
			firstSymbol++
			tempArr = nil
		}
	}
	d.MapOfBanner = tempChar
}
