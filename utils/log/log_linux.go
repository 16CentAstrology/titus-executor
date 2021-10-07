//go:build linux
// +build linux

package log

import (
	"github.com/wercker/journalhook"
)

// MaybeSetupLoggerIfOnJournaldAvailable sets up journald logging if the system
func MaybeSetupLoggerIfOnJournaldAvailable() {
	journalhook.Enable()
}
