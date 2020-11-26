package WireRouting

import "testing"

func TestIntersections(t *testing.T) {
	tests := []struct {
		first  string
		second string
		want   []node
	}{
		{first: "", second: "", want: []node{}},
		{first: "R8,U5,L5,D3", second: "", want: []node{}},
		{first: "", second: "U7,R6,D4,L4", want: []node{}},
		{first: "R8,U5,L5,D3", second: "U7,R6,D4,L4", want: []node{Node(3, 3), Node(6, 5)}},
		{first: "R1,U1,L1", second: "R1,U1,L1", want: []node{Node(1, 0), Node(1, 1), Node(0, 1)}},
	}
	for i, tc := range tests {
		one := Route(tc.first)
		two := Route(tc.second)
		got := one.Intersections(two)
		if len(got) != len(tc.want) {
			t.Errorf("Route(%v).Intersections(Route(%v)) want length %d, got %d (case %d)", tc.first, tc.second, len(tc.want), len(got), i)
		}
		for _, n := range tc.want {
			if !contains(got, n) {
				t.Errorf("Route(%v).Intersections(Route(%v)) want contains %v, does not (case %d)", tc.first, tc.second, n, i)
			}
		}
	}
}

func TestClosestPhysicalIntersections(t *testing.T) {
	tests := []struct {
		first  string
		second string
		want   int
	}{
		{first: "R75,D30,R83,U83,L12,D49,R71,U7,L72", second: "U62,R66,U55,R34,D71,R55,D58,R83", want: 159},
		{first: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", second: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", want: 135},
	}
	for i, tc := range tests {
		one := Route(tc.first)
		two := Route(tc.second)
		got := ClosestPhysical(one.Intersections(two))
		closest := got.ManhattanLength()
		if closest != tc.want {
			t.Errorf("Route(%v).Intersections(Route(%v)) want closest physical intersection %d away, got %d (case %d)", tc.first, tc.second, tc.want, closest, i)
		}
	}
}

func TestClosestRoutedIntersections(t *testing.T) {
	tests := []struct {
		first  string
		second string
		want   int
	}{
		{first: "R75,D30,R83,U83,L12,D49,R71,U7,L72", second: "U62,R66,U55,R34,D71,R55,D58,R83", want: 610},
		{first: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", second: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", want: 410},
	}
	for i, tc := range tests {
		one := Route(tc.first)
		two := Route(tc.second)
		got := ClosestRouted(one.Intersections(two))
		closest := got.RouteLength()
		if closest != tc.want {
			t.Errorf("Route(%v).Intersections(Route(%v)) want closest routed intersection %d away, got %d (case %d)", tc.first, tc.second, tc.want, closest, i)
		}
	}
}
