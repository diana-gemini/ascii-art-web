package internal

import (
	"fmt"
	"os"
)

// read text file and write to []byte
func ReadBannerFile(d *Ascii) error {
	var err error
	switch d.BannerName {
	case "standard":
		d.BannerFileName = "banner/standard.txt"
	case "shadow":
		d.BannerFileName = "banner/shadow.txt"
	case "thinkertoy":
		d.BannerFileName = "banner/thinkertoy.txt"
	}

	d.DataFromBannerFile, err = os.ReadFile(d.BannerFileName)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}
