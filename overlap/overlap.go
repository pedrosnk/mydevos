package overlap

type Line struct {
	x1, x2 int64
}

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

func OverlappedLines(line1, line2 Line) bool {
	return OverlappedCoords(line1.x1, line1.x2, line2.x1, line2.x2)
}

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
