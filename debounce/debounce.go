package debounce

import (
	"sync"
	"time"
)

//DefaultMaxDurationMagnification efault max duration magnification
var DefaultMaxDurationMagnification int64 = 2

//Debounce debounce struct
type Debounce struct {
	//Duration debounce duration
	Duration time.Duration
	//MaxDuration max lifetime debouce can live.
	MaxDuration time.Duration
	//Leading if the callback should be called before the duration
	Leading  bool
	lock     sync.Mutex
	deadline time.Time
	timer    *time.Timer
	d        chan int
	//Callback function
	Callback func()
}

func (d *Debounce) getTimer() *time.Timer {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.timer
}
func (d *Debounce) deleteTimer() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.timer = nil
}

//Discard Debounce timer
//Return true if reset success
func (d *Debounce) Discard() {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.timer == nil {
		return
	}
	d.timer.Stop()
	dc := d.d
	d.d = make(chan int)
	close(dc)
}

//Reset Debounce timer
//Return true if reset success
func (d *Debounce) Reset() bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.timer == nil {
		return false
	}
	if d.MaxDuration > 0 {
		d.deadline = time.Now().Add(d.MaxDuration)
	} else {
		d.deadline = time.Time{}
	}
	return d.timer.Reset(d.Duration)
}

//Exec call callback func.
//Return if the callback is executed immediately
func (d *Debounce) Exec() bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	var success = false
	var duration time.Duration
	if d.timer != nil {
		if !d.deadline.IsZero() {
			now := time.Now()
			duration = d.deadline.Sub(now)
			if duration > d.Duration {
				duration = d.Duration
			}
		} else {
			duration = d.Duration
		}
		success = d.timer.Reset(duration)
	}
	if success {
		return false
	}
	d.timer = time.NewTimer(d.Duration)
	if d.MaxDuration > 0 {
		d.deadline = time.Now().Add(d.MaxDuration)
	} else {
		d.deadline = time.Time{}
	}
	go func() {
		if d.Leading {
			d.Callback()
		}
		t := d.getTimer()
		if t != nil {
			select {

			case <-t.C:
				d.deleteTimer()
				if !d.Leading {
					d.Callback()
				}
			case <-d.d:
			}
		}
	}()
	if d.Leading {
		return true
	}
	return false

}

//WithLeading update debounce with given leading.
//Return debounce self.
func (d *Debounce) WithLeading(l bool) *Debounce {
	d.Leading = l
	return d
}

//WithMaxDuration update debounce with given max duration.
//Return debounce self.
func (d *Debounce) WithMaxDuration(duration time.Duration) *Debounce {
	d.MaxDuration = duration
	return d
}

//ExecFunc exec debounce
func (d *Debounce) ExecFunc() {
	d.Exec()
}

//New create new debounce with given duration and callback
func New(duration time.Duration, callback func()) *Debounce {
	d := &Debounce{}
	d.Duration = duration
	d.MaxDuration = time.Duration(DefaultMaxDurationMagnification) * duration
	d.d = make(chan int)
	d.Callback = callback
	return d
}
