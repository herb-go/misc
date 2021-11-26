package version

import "testing"

func TestDateVersion(t *testing.T) {
	v := &DateVersion{
		Major: 1,
		Year:  2021,
		Month: 11,
		Day:   26,
		Patch: 1,
		Build: "test build",
	}
	if v.VersionType() != "dateversion" {
		t.Fatal()
	}
	if v.MajorVersionCode() != "1" {
		t.Fatal()
	}
	if v.FullVersionCode() != "1.2021-11-26.1 test build" {
		t.Fatal(v.FullVersionCode())
	}
	v = &DateVersion{
		Major: 0,
		Year:  2021,
		Month: 1,
		Day:   3,
		Patch: 5,
		Build: "test build",
	}
	if v.MajorVersionCode() != "0" {
		t.Fatal()
	}
	if v.FullVersionCode() != "0.2021-01-03.5 test build" {
		t.Fatal(v.FullVersionCode())
	}
	v = &DateVersion{
		Major: 0,
		Year:  2021,
		Month: 1,
		Day:   3,
		Patch: 0,
		Build: "",
	}
	if v.FullVersionCode() != "0.2021-01-03" {
		t.Fatal(v.FullVersionCode())
	}
	v = &DateVersion{
		Major: -1,
		Year:  -2021,
		Month: -1,
		Day:   -3,
		Patch: -3,
		Build: "",
	}
	if v.FullVersionCode() != "-1.-2021--1--3.-3" {
		t.Fatal(v.FullVersionCode())
	}
	v = &DateVersion{
		Major: 1,
		Year:  2019,
		Month: 11,
		Day:   26,
		Patch: 1,
		Build: "test build",
	}
	v2 := &DateVersion{
		Major: 0,
		Year:  2020,
		Month: 1,
		Day:   30,
		Patch: 0,
		Build: "test build",
	}
	if !(v.MajorVersionWeight() > v2.MajorVersionWeight()) {
		t.Fatal()
	}
	if !(v.MinorVersionWeight() < v2.MinorVersionWeight()) {
		t.Fatal(v.MinorVersionWeight(), v2.MinorVersionWeight())
	}
}
