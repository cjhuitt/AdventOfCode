package pwparser

type span struct{ min, max int }

func parseRange(in string) (span, error) {
	return span{1, 3}, nil
}
