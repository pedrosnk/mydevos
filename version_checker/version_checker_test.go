package version_checker

import (
	"errors"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type CompareTest struct {
	v1, v2   string
	expected Comparation
}

type VCSuite struct{}

var _ = Suite(&VCSuite{})

func (s *VCSuite) TestParseSimpleString(c *C) {
	v, _ := Parse("1.2.0")

	c.Assert(&Version{major: 1, minor: 2, patch: 0}, DeepEquals, v)
}

func (s *VCSuite) TestParseWithPreValue(c *C) {
	v, _ := Parse("1.2.1-rc1")

	c.Assert(&Version{major: 1, minor: 2, patch: 1, pre: "rc1"}, DeepEquals, v)
}

func (s *VCSuite) TestCompareMajorVersions(c *C) {
	testCases := []CompareTest{
		{"1.0.0", "1.0.0", Eq},
		{"2.0.0", "1.0.0", Gt},
		{"1.0.0", "3.0.0", Lt},
	}

	for _, t := range testCases {
		result, _ := Compare(t.v1, t.v2)
		c.Assert(result, Equals, t.expected)
	}
}

func (s *VCSuite) TestCompareMinorVersions(c *C) {
	testCases := []CompareTest{
		{"1.1.0", "1.1.0", Eq},
		{"1.2.0", "1.1.0", Gt},
		{"1.1.0", "1.3.0", Lt},
	}

	for _, t := range testCases {
		result, _ := Compare(t.v1, t.v2)
		c.Assert(result, Equals, t.expected)
	}
}

func (s *VCSuite) TestComparePatchVersions(c *C) {
	testCases := []CompareTest{
		{"1.1.5", "1.1.5", Eq},
		{"1.1.10", "1.1.9", Gt},
		{"1.1.8", "1.1.12", Lt},
	}

	for _, t := range testCases {
		result, _ := Compare(t.v1, t.v2)
		c.Assert(result, Equals, t.expected)
	}
}

func (s *VCSuite) TestComparePreVersions(c *C) {
	testCases := []CompareTest{
		{"1.0.0-rc1", "1.0.0-rc1", Eq},
		{"1.0.0", "1.0.0-rc1", Gt},
		{"1.0.0-rc1", "1.0.0", Lt},
		{"1.0.0-rc1", "1.0.0-rc2", Lt},
		{"1.0.0-alpha", "1.0.0-beta", Lt},
		{"1.0.0-beta", "1.0.0-alpha", Gt},
	}

	for _, t := range testCases {
		result, _ := Compare(t.v1, t.v2)
		c.Assert(result, Equals, t.expected)
	}
}

func (s *VCSuite) TestCompareReturnsErrorOnInvalidString(c *C) {
	testCases := []CompareTest{
		{"1.0", "1.0.0", Invalid},
		{"1.0.0", "2", Invalid},
		{"1.0.0.0", "1.0.0", Invalid},
		{"a.0.0", "1.0.0", Invalid},
		{"1.b.0", "1.0.0", Invalid},
		{"1.0.c", "1.0.0", Invalid},
	}

	errs := []error{
		errors.New("Invalid version 1.0"),
		errors.New("Invalid version 2"),
		errors.New("Invalid version 1.0.0.0"),
		errors.New("Invalid major value a"),
		errors.New("Invalid minor value b"),
		errors.New("Invalid patch value c"),
	}

	for i, t := range testCases {
		result, err := Compare(t.v1, t.v2)
		c.Assert(err, DeepEquals, errs[i])
		c.Assert(result, Equals, t.expected)
	}
}
