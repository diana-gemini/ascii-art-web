package internal

type Ascii struct {
	BannerName         string
	BannerFileName     string
	DataFromBannerFile []byte
	MapOfBanner        map[rune][]string
	InputText          string
	InputSlice         []string
	OutputText         string
}

type Err struct {
	StatusText string
	StatusCode int
}
