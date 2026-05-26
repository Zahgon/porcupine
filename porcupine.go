package porcupine

import "time"

// CheckOperations checks whether a history is linearizable.
func CheckOperations(model Model, history []Operation) bool {
	_ = "STUB: not implemented"
	return false
}

// CheckOperationsTimeout checks whether a history is linearizable, with a
// timeout.
//
// A timeout of 0 is interpreted as an unlimited timeout.
func CheckOperationsTimeout(model Model, history []Operation, timeout time.Duration) CheckResult {
	_ = "STUB: not implemented"
	return *new(CheckResult)
}

// CheckOperationsVerbose checks whether a history is linearizable while
// computing data that can be used to visualize the history and linearization.
//
// The returned LinearizationInfo can be used with [Visualize].
func CheckOperationsVerbose(model Model, history []Operation, timeout time.Duration) (CheckResult, LinearizationInfo) {
	_ = "STUB: not implemented"
	return *new(CheckResult), *new(LinearizationInfo)
}

// CheckEvents checks whether a history is linearizable.
func CheckEvents(model Model, history []Event) bool { _ = "STUB: not implemented"; return false }

// CheckEventsTimeout checks whether a history is linearizable, with a timeout.
//
// A timeout of 0 is interpreted as an unlimited timeout.
func CheckEventsTimeout(model Model, history []Event, timeout time.Duration) CheckResult {
	_ = "STUB: not implemented"
	return *new(CheckResult)
}

// CheckEventsVerbose checks whether a history is linearizable while computing
// data that can be used to visualize the history and linearization.
//
// The returned LinearizationInfo can be used with [Visualize].
func CheckEventsVerbose(model Model, history []Event, timeout time.Duration) (CheckResult, LinearizationInfo) {
	_ = "STUB: not implemented"
	return *new(CheckResult), *new(LinearizationInfo)
}
