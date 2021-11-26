package version

import "testing"

func TestSemver(t *testing.T) {
	v := &Semver{
		Major: 1,
		Minor: 1,
		Patch: 1,
		Build: "test build",
	}
	if v.VersionType() != "semver" {
		t.Fatal()
	}
	if v.MajorVersionCode() != "1" {
		t.Fatal()
	}
	if v.FullVersionCode() != "1.1.1 test build" {
		t.Fatal(v.FullVersionCode())
	}
	v = &Semver{
		Major: 0,
		Minor: 1,
		Patch: 1,
		Build: "test build",
	}
	if v.MajorVersionCode() != "0" {
		t.Fatal()
	}
	if v.FullVersionCode() != "0.1.1 test build" {
		t.Fatal(v.FullVersionCode())
	}
	v = &Semver{
		Major: 0,
		Minor: 1,
		Patch: 0,
		Build: "",
	}
	if v.FullVersionCode() != "0.1" {
		t.Fatal(v.FullVersionCode())
	}
	v = &Semver{
		Major: -1,
		Minor: -1,
		Patch: -3,
		Build: "",
	}
	if v.FullVersionCode() != "-1.-1.-3" {
		t.Fatal(v.FullVersionCode())
	}
	v = &Semver{
		Major: 2,
		Minor: 1,
		Patch: 0,
		Build: "test build",
	}
	v2 := &Semver{
		Major: 1,
		Minor: 2,
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
