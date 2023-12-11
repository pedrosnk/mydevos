package version_checker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// User semver to parse, and compare software version

type Version struct {
	major, minor, patch int64
	pre                 string
}

type Comparation int64

const (
	Lt      Comparation = -1
	Eq      Comparation = 0
	Gt      Comparation = 1
	Invalid Comparation = 2
)

func Compare(str1, str2 string) (Comparation, error) {
	v1, err := Parse(str1)
	if err != nil {
		return Invalid, err
	}

	v2, err := Parse(str2)
	if err != nil {
		return Invalid, err
	}

	if v1.major > v2.major {
		return Gt, nil
	} else if v1.major < v2.major {
		return Lt, nil
	} else if v1.minor > v2.minor {
		return Gt, nil
	} else if v1.minor < v2.minor {
		return Lt, nil
	} else if v1.patch > v2.patch {
		return Gt, nil
	} else if v1.patch < v2.patch {
		return Lt, nil
	} else {
		// If value is equal, check pre values
		comparation := strings.Compare(v1.pre, v2.pre)
		if v1.pre == "" || v2.pre == "" {
			return Comparation(-comparation), nil
		} else {
			return Comparation(comparation), nil
		}
	}
}

func Parse(str string) (*Version, error) {
	version := &Version{}
	versions := strings.SplitN(str, "-", 2)
	if len(versions) == 2 {
		version.pre = versions[1]
	}
	versions = strings.Split(versions[0], ".")

	if len(versions) != 3 {
		return nil, errors.New(fmt.Sprintf("Invalid version %s", str))
	}

	number, err := strconv.ParseInt(versions[0], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid major value %s", versions[0]))
	}
	version.major = number

	number, err = strconv.ParseInt(versions[1], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid minor value %s", versions[1]))
	}
	version.minor = number

	number, err = strconv.ParseInt(versions[2], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid patch value %s", versions[2]))
	}
	version.patch = number

	return version, nil
}
