package tickets

type constraint struct {
	min, max int
}

func parseConstraint(in string) constraint {
	return constraint{-1, -1}
}
