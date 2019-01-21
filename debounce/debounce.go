package debounce

import (
	"sync"
	"time"
)

type Debounce struct {
	Duration    time.Duration
	MaxDuration time.Duration
	Leading     bool
	lock        sync.Mutex
	deadline    time.Time
	timer       *time.Timer
	Callback    func()
}

func (d *Debounce) Exec() bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.timer == nil {
		d.timer = time.NewTimer(d.Duration)
		d.deadline = time.Now().Add(d.MaxDuration)
		go func() {
			if d.Leading {
				d.Callback()
			}
			<-d.timer.C
			d.timer = nil
			if !d.Leading {
				d.Callback()
			}
		}()
		if d.Leading {
			return true
		}
		return false
	}
	now := time.Now()
	duration := d.deadline.Sub(now)
	if duration > d.Duration {
		duration = d.Duration
	}
	success := d.timer.Reset(duration)
	if success == false {
		go func() {
			d.Exec()
		}()
		if d.Leading {
			return true
		}
		return false
	}
	return false
}
func New(duration time.Duration, callback func()) *Debounce {
	d := &Debounce{}
	d.Duration = duration
	d.MaxDuration = 2 * duration
	d.Callback = callback
	return d
}
