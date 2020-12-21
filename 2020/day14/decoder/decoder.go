package decoder

type mask struct {
	allowed    uint64
	predefined uint64
}

func parseMask(in string) mask {
	var allowed, predef uint64
	for _, c := range in {
		allowed = allowed << 1
		predef = predef << 1
		switch c {
		case 'X':
			allowed += 1
		case '1':
			allowed += 0
			predef += 1
		case '0':
			allowed += 0
		}
	}

	return mask{allowed, predef}
}
