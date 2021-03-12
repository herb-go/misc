package waitingmap

import (
	"testing"
	"time"
)

func TestOnceMap(t *testing.T) {
	var count = 0
	var okCount = 0
	f := func() {
		count++
		time.Sleep(100 * time.Millisecond)
	}
	m := NewOnceMap()
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	time.Sleep(300 * time.Millisecond)
	if count != 1 || okCount != 1 {
		t.Fatal(count, okCount)
	}
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	time.Sleep(300 * time.Millisecond)
	if count != 2 || okCount != 2 {
		t.Fatal(count, okCount)
	}
	go func() {
		ok := m.Do("test", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	go func() {
		ok := m.Do("test2", f)
		if ok {
			okCount = okCount + 1
		}
	}()
	time.Sleep(300 * time.Millisecond)
	if count != 4 || okCount != 4 {
		t.Fatal(count, okCount)
	}
}
