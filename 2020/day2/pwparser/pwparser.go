package pwparser

type span struct{ min, max int }

func parseSpan(in string) (span, error) {
	return span{1, 3}, nil
}
