package debounce

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	var value = 1
	var add = func() {
		value++
	}
	var result bool
	var testDuration = time.Millisecond
	var TestDebounce = New(2*testDuration, add)
	if TestDebounce.MaxDuration != 2*2*testDuration {
		t.Fatal(TestDebounce.MaxDuration)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 1 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 1 {
		t.Fatal(value)
	}
	time.Sleep(3 * time.Millisecond)
	if value != 2 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 2 {
		t.Fatal(value)
	}
	time.Sleep(3 * time.Millisecond)
	if value != 3 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 3 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 3 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 3 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 3 {
		t.Fatal(value)
	}
	time.Sleep(2 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 4 {
		t.Fatal(value)
	}
}

func TestDebounceLeading(t *testing.T) {
	var value = 1
	var add = func() {
		value++
	}
	var result bool
	var testDuration = time.Millisecond
	var TestDebounce = New(2*testDuration, add)
	TestDebounce.Leading = true
	if TestDebounce.MaxDuration != 2*2*testDuration {
		t.Fatal(TestDebounce.MaxDuration)
	}
	result = TestDebounce.Exec()
	if !result {
		t.Fatal(result)
	}
	time.Sleep(time.Microsecond)
	if value != 2 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 2 {
		t.Fatal(value)
	}
	time.Sleep(3 * time.Millisecond)
	if value != 2 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if !result {
		t.Fatal(result)
	}
	time.Sleep(time.Microsecond)
	if value != 3 {
		t.Fatal(value)
	}
	time.Sleep(3 * time.Millisecond)
	if value != 3 {
		t.Fatal(value)
	}
	result = TestDebounce.Exec()
	if !result {
		t.Fatal(result)
	}
	time.Sleep(time.Microsecond)
	if value != 4 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 4 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 4 {
		t.Fatal(value)
	}
	time.Sleep(1 * time.Millisecond)
	result = TestDebounce.Exec()
	if result {
		t.Fatal(result)
	}
	if value != 4 {
		t.Fatal(value)
	}
	time.Sleep(2 * time.Millisecond)
	result = TestDebounce.Exec()
	if !result {
		t.Fatal(result)
	}
	time.Sleep(time.Microsecond)

	if value != 5 {
		t.Fatal(value)
	}
}
