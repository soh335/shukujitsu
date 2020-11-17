package shukujitsu

import (
	"time"
)

// IsShukujitsu datermin shukujitsu.
// This function don't care timezone of arguments time.
// So you need to care your timezone.
func IsShukujitsu(t time.Time) bool {
	key := t.Format("2006/1/2")
	if _, ok := dates[key]; ok {
		return true
	}
	return false
}
