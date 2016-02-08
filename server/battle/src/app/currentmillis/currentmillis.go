package currentmillis

import "time"

// StubNow For Test.
var StubNow func() int64

// Second.
const Second int64 = 1000

// Minute.
const Minute int64 = 60 * Second

// Now returns the number of the current time milliseconds elapsed since UNIX epoch time.
func Now() int64 {
	if StubNow != nil {
		return StubNow()
	}
	return time.Now().UnixNano() / 1000000
}

// ToMilliseconds returns the number of the given time milliseconds elapsed since UNIX epoch time.
func ToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// Time returns the local Time corresponding to the given currentMillis,
func Time(currentMillis int64) time.Time {
	sec := currentMillis / 1000
	nsec := (currentMillis % 1000) * 1000000
	return time.Unix(sec, nsec)
}
