package internal

func CollectOutputText(d *Ascii) {
	var out string
	var countOfOnlyNewLine int

	for _, val := range d.InputSlice {
		if val == "" {
			countOfOnlyNewLine++
		}
	}

	if countOfOnlyNewLine == len(d.InputSlice) {
		for i := 0; i < len(d.InputSlice)-1; i++ {
			out += "\n"
		}
		d.OutputText = out
		return
	}

	for _, val := range d.InputSlice {
		if len(val) != 0 {
			for i := 0; i < 8; i++ {
				for _, symb := range val {
					out += d.MapOfBanner[symb][i]
				}
				out += "\n"
			}
		} else if len(val) == 0 {
			out += "\n"
		}
	}
	d.OutputText = out
}
