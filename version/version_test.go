package version

import "testing"

func TestVersion(t *testing.T) {
	var v = &Semver{
		Major: 1,
		Minor: 1,
	}
	var v2 = &Semver{
		Major: 1,
		Minor: 1,
	}
	if CompareMajorVersion(v, v2) != 0 {
		t.Fatal()
	}
	if CompareMinorVersion(v, v2) != 0 {
		t.Fatal()
	}
	if CompareVersion(v, v2) != 0 {
		t.Fatal()
	}
	v2 = &Semver{
		Major: 2,
		Minor: 2,
	}
	if CompareMajorVersion(v, v2) != -1 {
		t.Fatal()
	}
	if CompareMinorVersion(v, v2) != -1 {
		t.Fatal()
	}
	if CompareVersion(v, v2) != -1 {
		t.Fatal()
	}
	v2 = &Semver{
		Major: 0,
		Minor: 0,
	}
	if CompareMajorVersion(v, v2) != 1 {
		t.Fatal()
	}
	if CompareMinorVersion(v, v2) != 1 {
		t.Fatal()
	}
	if CompareVersion(v, v2) != 1 {
		t.Fatal()
	}
	v2 = &Semver{
		Major: 0,
		Minor: 2,
	}
	if CompareMajorVersion(v, v2) != 1 {
		t.Fatal()
	}
	if CompareMinorVersion(v, v2) != -1 {
		t.Fatal()
	}
	if CompareVersion(v, v2) != 1 {
		t.Fatal()
	}
	v2 = &Semver{
		Major: 1,
		Minor: 2,
	}
	if CompareMajorVersion(v, v2) != 0 {
		t.Fatal()
	}
	if CompareMinorVersion(v, v2) != -1 {
		t.Fatal()
	}
	if CompareVersion(v, v2) != -1 {
		t.Fatal()
	}
}

func TestError(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil || r.(error) != ErrVersionTypeNotMatch {
			t.Fatal()
		}
	}()
	CompareMajorVersion(&Semver{}, &DateVersion{})
}
func TestError2(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil || r.(error) != ErrVersionTypeNotMatch {
			t.Fatal()
		}
	}()
	CompareMinorVersion(&Semver{}, &DateVersion{})
}
