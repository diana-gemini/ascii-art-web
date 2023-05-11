package internal

import (
	"errors"
	"regexp"
	"strings"
)

func CheckInputText(d *Ascii) error {
	var err error

	for _, val := range d.InputText {
		if !(val >= 0 && val <= '~') {
			err = errors.New("not ASCII characters")
			return err
		}
	}

	if len(d.InputText) > 400 {
		err = errors.New("large input text")
		return err
	}

	d.InputText = strings.ReplaceAll(d.InputText, "\\n", "\n")

	// if len(d.InputText) == 0 {
	// 	err = errors.New("give some text")
	// 	return err
	// }

	symbol := regexp.MustCompile(`\r`)
	strText := symbol.ReplaceAllString(string(d.InputText), "")
	d.InputSlice = strings.Split(strText, "\n")

	return err
}
