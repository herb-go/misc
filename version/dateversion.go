package version

import "strconv"

type DateVersion struct {
	Major int
	Year  int
	Month int
	Day   int
	Patch int
	Build string
}

//MajorVersionCode code for major version.
//For example "V7"
func (v *DateVersion) MajorVersionCode() string {
	return strconv.Itoa(v.Major)
}

//MajorVersionWeight weight for major version
//Should only used when compare two version
func (v *DateVersion) MajorVersionWeight() int64 {
	return int64(v.Major)
}

//FullVersionCode full code for version
//For example "V7.23.44 nightly build"
func (v *DateVersion) FullVersionCode() string {

	code := strconv.Itoa(v.Major)
	code = code + "."
	code = code + strconv.Itoa(v.Year)
	code = code + "-"
	if v.Month < 10 && v.Month >= 0 {
		code = code + "0" + strconv.Itoa(v.Month)
	} else {
		code = code + strconv.Itoa(v.Month)
	}
	code = code + "-"
	if v.Day < 10 && v.Day >= 0 {
		code = code + "0" + strconv.Itoa(v.Day)
	} else {
		code = code + strconv.Itoa(v.Day)
	}
	if v.Patch != 0 {
		code = code + "."
		code = code + strconv.Itoa(v.Patch)
	}
	if v.Build != "" {
		code = code + " " + v.Build
	}
	return code
}

//MinorVersionWeight weight for MinorVersionWeight
//Should only used when compare two version
func (v *DateVersion) MinorVersionWeight() int64 {
	var weight int64
	weight = weight + int64(uint16(v.Year))
	weight = weight << 8
	weight = weight + int64(uint8(v.Month))
	weight = weight << 8
	weight = weight + int64(uint8(v.Day))
	weight = weight << 16
	weight = weight + int64(uint16(v.Patch))
	return weight
}

//VersionType version type
func (v *DateVersion) VersionType() string {
	return "dateversion"
}
