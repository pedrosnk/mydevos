package overlap

type Line struct {
	X1, X2 int64
}

// Given two line coords, returns true if the lines overlaps
// with each other.
//
// Example 1: Line(1, 3) and Line(4, 6) doesn't overlap.
// OverlappedCoords(1, 3, 4, 6) => false
//
// Example1: Line(1, 3) and Line(2, 6) overlap.
// OverlappedCoords(1, 3, 2, 6) => true
func OverlappedCoords(x1, x2, x3, x4 int64) bool {
	// To simplify further logic, swap numbers if order is desc
	x1, x2 = swapIfNeeded(x1, x2)
	x3, x4 = swapIfNeeded(x3, x4)

	if x2 <= x3 || x1 >= x4 {
		return false
	} else {
		return true
	}
}

// Same as OverlappedCoords, but expects line struct to be passed on.
//
// Example 1:
// l1 = Line{X1: 1, X2: 3}
// l2 = Line{X1: 4, X2: 6}
// OverlappedLines(l1, l2) => false
//
// Example 2:
// l1 = Line{X1: 1, X2: 3}
// l2 = Line{X1: 2, X2: 6}
// OverlappedLines(l1, l2) => true
func OverlappedLines(line1, line2 Line) bool {
	return OverlappedCoords(line1.X1, line1.X2, line2.X1, line2.X2)
}

// Same as OverlappedCoords, but it is binded to a Line Struct
//
// Example 1:
// l = Line{X1: 1, X2: 3}
// l.OverlapsWith(Line{X1: 4, X2: 6})  => false
// l.OverlapsWith(Line{X1: 2, X2: 6})  => true
func (l *Line) OverlapsWith(other Line) bool {
	return OverlappedLines(*l, other)
}

func swapIfNeeded(x1, x2 int64) (int64, int64) {
	if x1 > x2 {
		return x2, x1
	} else {
		return x1, x2
	}
}
