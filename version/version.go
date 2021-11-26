package version

// Version interface
type Version interface {
	//MajorVersionCode code for major version.
	//For example "V7"
	MajorVersionCode() string
	//MajorVersionWeight weight for major version
	//Should only used when compare two version
	MajorVersionWeight() int64
	//FullVersionCode full code for version
	//For example "V7.23.44 nightly build"
	FullVersionCode() string
	//MinorVersionWeight weight for MinorVersionWeight
	//Should only used when compare two version
	MinorVersionWeight() int64
	//VersionType version type
	VersionType() string
}

//CompareMajorVersion compare two version for major version
//Return 1 if v1 < v2,0 if equle,1 if v1>v2
//Panic if v1 and v2 has different type
func CompareMajorVersion(v1, v2 Version) int {
	if v1.VersionType() != v2.VersionType() {
		panic(ErrVersionTypeNotMatch)
	}
	d := v1.MajorVersionWeight() - v2.MajorVersionWeight()
	if d < 0 {
		return -1
	} else if d == 0 {
		return 0
	}
	return 1
}

//CompareMinorVersion compare two version for minor version
//Return 1 if v1 < v2,0 if equle,1 if v1>v2
//Panic if v1 and v2 has different type
func CompareMinorVersion(v1, v2 Version) int {
	if v1.VersionType() != v2.VersionType() {
		panic(ErrVersionTypeNotMatch)
	}
	d := v1.MinorVersionWeight() - v2.MinorVersionWeight()
	if d < 0 {
		return -1
	} else if d == 0 {
		return 0
	}
	return 1
}

//CompareVersion compare two version
//Return 1 if v1 < v2,0 if equle,1 if v1>v2
//Panic if v1 and v2 has different type
func CompareVersion(v1, v2 Version) int {
	d := CompareMajorVersion(v1, v2)
	if d != 0 {
		return d
	}
	return CompareMinorVersion(v1, v2)
}
