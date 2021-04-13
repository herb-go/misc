package waitingmap

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestLockerMap(t *testing.T) {
	var result = []string{}
	lm := NewLockerMap()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		time.Sleep(0)
		newlock := lm.Lock("test1")
		if !newlock {
			panic(newlock)
		}
		result = append(result, "locktest1")
		time.Sleep(2 * time.Millisecond)
		free := lm.Unlock("test1")
		if free {
			panic(free)
		}
		result = append(result, "unlocktest1")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		newlock := lm.Lock("test1")
		if newlock {
			panic(newlock)
		}
		result = append(result, "locktest1*")
		time.Sleep(2 * time.Millisecond)
		free := lm.Unlock("test1")
		if !free {
			panic(free)
		}
		result = append(result, "unlocktest1*")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Millisecond)
		newlock := lm.Lock("test2")
		if !newlock {
			panic(newlock)
		}
		result = append(result, "locktest2")
		time.Sleep(3 * time.Millisecond)
		free := lm.Unlock("test2")
		if !free {
			panic(free)
		}
		result = append(result, "unlocktest2")
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Millisecond)
		newlock := lm.RLock("test1")
		if newlock {
			panic(newlock)
		}
		result = append(result, "rlocktest1")
		time.Sleep(4 * time.Millisecond)
		free := lm.RUnlock("test1")
		if free {
			panic(free)
		}
		result = append(result, "runlocktest1")
		wg.Done()
	}()
	wg.Wait()
	if strings.Join(result, ",") != "locktest1,locktest2,unlocktest1,rlocktest1,unlocktest2,runlocktest1,locktest1*,unlocktest1*" {
		t.Fatal(strings.Join(result, ","))
	}
}
func TestLockerMapMethods(t *testing.T) {
	lm := NewLockerMap()
	lm.RLock("test1")
	lm.RUnlock("test1")
	if len(lm.store) != 0 {
		t.Fatal()
	}
	l := lm.Locker("test")
	rl := lm.RLocker("test")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		rl.Lock()
		rl.Unlock()
		wg.Done()
	}()
	go func() {
		l.Lock()
		wg.Done()
		if len(lm.store) != 1 {
			panic("Err")
		}
		l.Unlock()
	}()
	wg.Wait()
	if len(lm.store) != 0 {
		t.Fatal()
	}
}
