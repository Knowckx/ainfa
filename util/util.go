package util

func ShortStr(in string) string {
	leng := 7 // length of short git commitID
	return ShortStrLen(in, leng)
}

func ShortStrLen(in string, leng int) string {
	out := in
	if len(out) > leng {
		out = out[0:leng]
	}
	return out
}

// Compress Spaces and Tap to one space
func CompressSpaces(in string) string {
	out := ""
	var added int
	for _, key := range in {
		if key == ' ' || key == '	' {
			if added == 0 {
				out += " "
				added = 1
			} else {
				added++
			}
		} else {
			out += string(key)
			added = 0
		}
	}
	return out
}
