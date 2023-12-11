package overlap

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type lineTest struct {
	x1, x2, x3, x4 int64
	expected       bool
}

var testTable = []lineTest{
	{x1: 1, x2: 3, x3: 4, x4: 6, expected: false},
	{x1: 5, x2: 2, x3: 5, x4: 6, expected: false},
	{x1: 5, x2: 2, x3: 5, x4: 6, expected: false},
	{x1: -10, x2: -5, x3: -3, x4: 1, expected: false},
	{x1: 4, x2: 8, x3: 1, x4: 3, expected: false},
	{x1: 3, x2: 5, x3: 3, x4: 6, expected: true},
	{x1: 3, x2: 10, x3: 9, x4: 12, expected: true},
	{x1: 9, x2: 12, x3: 10, x4: 3, expected: true},
	{x1: 12, x2: 4, x3: 3, x4: 10, expected: true},
}

type OverlapSuite struct{}

var _ = Suite(&OverlapSuite{})

func (s *OverlapSuite) TestOverlap(c *C) {
	for _, t := range testTable {
		c.Assert(OverlappedCoords(t.x1, t.x2, t.x3, t.x4), Equals, t.expected)
	}
}

//func (s *OverlapSuite) TestCallingOverlaps(c *C) {
//  for _, t := range testTable {
//    l1 := Line{x1: t.x1, x2: t.x2}
//    l2 := Line{x1: t.x3, x2: t.x4}
//
//    c.Assert(l1.OverlapsWith(l2), Equals, t.expected)
//  }
//}
