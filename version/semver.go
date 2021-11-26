package version

import "strconv"

type Semver struct {
	Major int
	Minor int
	Patch int
	Build string
}

//MajorVersionCode code for major version.
//For example "V7"
func (s *Semver) MajorVersionCode() string {
	return strconv.Itoa(s.Major)
}

//MajorVersionWeight weight for major version
//Should only used when compare two version
func (s *Semver) MajorVersionWeight() int64 {
	return int64(s.Major)
}

//FullVersionCode full code for version
//For example "V7.23.44 nightly build"
func (s *Semver) FullVersionCode() string {

	code := strconv.Itoa(s.Major)
	code = code + "."
	code = code + strconv.Itoa(s.Minor)
	if s.Patch != 0 {
		code = code + "."
		code = code + strconv.Itoa(s.Patch)
	}
	if s.Build != "" {
		code = code + " " + s.Build
	}
	return code
}

//MinorVersionWeight weight for MinorVersionWeight
//Should only used when compare two version
func (s *Semver) MinorVersionWeight() int64 {
	var weight int64
	weight = weight + int64(uint32(s.Minor))
	weight = weight << 32
	weight = weight + int64(uint32(s.Patch))
	return weight
}

//VersionType version type
func (s *Semver) VersionType() string {
	return "semver"
}
