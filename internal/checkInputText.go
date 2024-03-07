package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func CheckInputText(d *Ascii) (err error) {
	defer func() {
		panicValue := recover()
		if panicValue != nil {
			fmt.Println(panicValue)
		}
	}()

	for _, val := range d.InputText {
		if !(val >= 0 && val <= '~') {
			err = errors.New("not ASCII characters")
			panic("not ASCII characters")
		}
	}

	if len(d.InputText) > 400 {
		err = errors.New("large input text")
		panic("large input text")
	}

	d.InputText = strings.ReplaceAll(d.InputText, "\\n", "\n")

	symbol := regexp.MustCompile(`\r`)
	strText := symbol.ReplaceAllString(string(d.InputText), "")
	d.InputSlice = strings.Split(strText, "\n")

	return err
}
