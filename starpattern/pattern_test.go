package starpattern

import "testing"

func equalFound(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
func TestStar(t *testing.T) {
	var ok bool
	var found []string
	ok, found = New("*").Find("")
	if !ok || !equalFound(found, []string{""}) {
		t.Fatal(ok, found)
	}
	ok, found = New("").Find("")
	if !ok || !equalFound(found, []string{}) {
		t.Fatal(ok, found)
	}

	ok, found = New("abc").Find("abc")
	if !ok || !equalFound(found, []string{"abc"}) {
		t.Fatal(ok, found)
	}
	ok, found = New("abc").Find("ab")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = New("abc").Find("abcd")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = New("*abc").Find("1abc")
	if !ok || !equalFound(found, []string{"1abc", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = New("*ab").Find("abc")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = New("*ab").Find("abcd")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = New("abc*").Find("abc")
	if !ok || !equalFound(found, []string{"abc", ""}) {
		t.Fatal(ok, found)
	}
	ok, found = New("abc*").Find("abc1")
	if !ok || !equalFound(found, []string{"abc1", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = New("abc**").Find("abc1")
	if !ok || !equalFound(found, []string{"abc1", "", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = New("abc*def*").Find("abcdef")
	if !ok || !equalFound(found, []string{"abcdef", "", ""}) {
		t.Fatal(ok, found)
	}
}

func TestException(t *testing.T) {
	var ok bool
	var found []string
	var opt = &Options{
		Wildcard:  '+',
		Exception: []rune("./"),
	}
	ok, found = opt.New("+.abc/123").Find("www.abc/123")
	if !ok || !equalFound(found, []string{"www.abc/123", "www"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("+.abc/").Find("www.abc/123")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("+.abc/1234").Find("www.abc/123")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
}

func TestAvaliable(t *testing.T) {
	var ok bool
	var found []string
	var opt = &Options{
		Wildcard:  '+',
		Avaliable: []rune("wabc123"),
	}
	ok, found = opt.New("+.abc/123").Find("www.abc/123")
	if !ok || !equalFound(found, []string{"www.abc/123", "www"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("+.abc/").Find("www.abc/123")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("+.abc/1234").Find("www.abc/123")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
}

func TestNotEmpty(t *testing.T) {
	var ok bool
	var found []string
	opt := &Options{
		Wildcard: '*',
		NotEmpty: true,
	}
	ok, found = opt.New("*").Find("")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("").Find("")
	if !ok || !equalFound(found, []string{}) {
		t.Fatal(ok, found)
	}

	ok, found = opt.New("abc").Find("abc")
	if !ok || !equalFound(found, []string{"abc"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc").Find("ab")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc").Find("abcd")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("*abc").Find("1abc")
	if !ok || !equalFound(found, []string{"1abc", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("*ab").Find("abc")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("*ab").Find("abcd")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc*").Find("abc")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc*").Find("abc1")
	if !ok || !equalFound(found, []string{"abc1", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc**").Find("abc1")
	if !ok || !equalFound(found, []string{"abc1", "", "1"}) {
		t.Fatal(ok, found)
	}
	ok, found = opt.New("abc*def*").Find("abcdef")
	if ok || found != nil {
		t.Fatal(ok, found)
	}
}

func TestEqual(t *testing.T) {
	if Star.equalRunes([]rune("abc"), []rune("abc1")) {
		t.Fatal()
	}
	if Star.equalRunes([]rune("abc"), []rune("Abc")) {
		t.Fatal()
	}
	opt := &Options{
		Wildcard:   '*',
		IgnoreCase: true,
	}
	if !opt.equalRunes([]rune("abc"), []rune("Abc")) {
		t.Fatal()
	}
}

func TestMatch(t *testing.T) {
	var ok bool
	ok = New("*").Match("")
	if !ok {
		t.Fatal(ok)
	}
}
