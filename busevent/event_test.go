package busevent

import "testing"

func newTestHandler(val string) func(data interface{}) {
	return func(data interface{}) {
		var s = data.(*string)
		*s = *s + val
	}
}

func TestEmpty(t *testing.T) {
	var result string
	var e Event
	e.Raise(&result)
	if result != "" {
		t.Fatal()
	}
	e.Bind(newTestHandler("a"))
	e.Raise(&result)
	if result != "a" {
		t.Fatal()
	}
	e.BindAs("b", newTestHandler("b"))
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove("noexist")
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove(nil)
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove("b")
	result = ""
	e.Raise(&result)
	if result != "a" {
		t.Fatal(result)
	}
	e.BindAs("b", newTestHandler("b"))
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Flush()
	result = ""
	result = ""
	e.Raise(&result)
	if result != "" {
		t.Fatal()
	}
}

func TestEvent(t *testing.T) {
	var result string
	var e = New()
	e.Raise(&result)
	if result != "" {
		t.Fatal()
	}
	e.Bind(newTestHandler("a"))
	e.Raise(&result)
	if result != "a" {
		t.Fatal()
	}
	e.BindAs("b", newTestHandler("b"))
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove("noexist")
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove(nil)
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Remove("b")
	result = ""
	e.Raise(&result)
	if result != "a" {
		t.Fatal(result)
	}
	e.BindAs("b", newTestHandler("b"))
	result = ""
	e.Raise(&result)
	if result != "ab" {
		t.Fatal()
	}
	e.Flush()
	result = ""
	result = ""
	e.Raise(&result)
	if result != "" {
		t.Fatal()
	}
}
