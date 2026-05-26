package porcupine

import (
	"time"
)

type entryKind bool

const (
	callEntry   entryKind = false
	returnEntry entryKind = true
)

type entry struct {
	kind     entryKind
	value    interface{}
	id       int
	time     int64
	clientId int
	metadata interface{}
}

type LinearizationInfo struct {
	history               [][]entry // for each partition, a list of entries
	partialLinearizations [][][]int // for each partition, a set of histories (list of ids)
	annotations           []Annotation
}

// PartialLinearizations returns partial linearizations found during the
// linearizability check, as sets of operation IDs.
//
// For each partition, it returns a set of possible linearization histories,
// where each history is represented as a sequence of operation IDs. If the
// history is linearizable, this will contain a complete linearization. If not
// linearizable, it contains the maximal partial linearizations found.
func (li *LinearizationInfo) PartialLinearizations() [][][]int {
	_ = "STUB: not implemented"
	return nil
}

// PartialLinearizationsOperations returns partial linearizations found during
// the linearizability check, as sets of sequences of [Operation].
//
// For each partition, it returns a set of possible linearization histories,
// where each history is represented as a sequence of [Operation]. If the
// history is linearizable, this will contain a complete linearization. If not
// linearizable, it contains the maximal partial linearizations found.
func (li *LinearizationInfo) PartialLinearizationsOperations() [][][]Operation {
	_ = "STUB: not implemented"
	return nil
}

// reconstruct operations based on entries

// this should never happen, because the LinearizationInfo
// object should always contain valid partial linearizations,
// where there is a return for every call

// prefer return metadata over call metadata

// this should never happen, because the LinearizationInfo
// object should always contain valid partial
// linearizations, where every ID in the partial
// linearization is in the history

type byTime []entry

func (a byTime) Len() int { _ = "STUB: not implemented"; return 0 }

func (a byTime) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (a byTime) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// if the timestamps are the same, we need to make sure we order calls
// before returns

func makeEntries(history []Operation) []entry { _ = "STUB: not implemented"; return nil }

type node struct {
	value interface{}
	match *node // call if match is nil, otherwise return
	id    int
	next  *node
	prev  *node
}

func insertBefore(n *node, mark *node) *node { _ = "STUB: not implemented"; return nil }

func length(n *node) int { _ = "STUB: not implemented"; return 0 }

func renumber(events []Event) []Event { _ = "STUB: not implemented"; return nil }

// renumbering

func convertEntries(events []Event) []entry { _ = "STUB: not implemented"; return nil }

// use index as "time"

func makeLinkedEntries(entries []entry) *node { _ = "STUB: not implemented"; return nil }

type cacheEntry struct {
	linearized bitset
	state      interface{}
}

func cacheContains(model Model, cache map[uint64][]cacheEntry, entry cacheEntry) bool {
	_ = "STUB: not implemented"
	return false
}

type callsEntry struct {
	entry *node
	state interface{}
}

func lift(entry *node) { _ = "STUB: not implemented"; return }

func unlift(entry *node) { _ = "STUB: not implemented"; return }

func checkSingle(model Model, history []entry, computePartial bool, kill *int32) (bool, []*[]int) {
	_ = "STUB: not implemented"
	return false, nil
}

// map from hash to cache entry

// longest linearizable prefix that includes the given entry

// the return entry

// longest

// create seq lazily

// longest linearization is the complete linearization, which is calls

func fillDefault(model Model) Model { _ = "STUB: not implemented"; return *new(Model) }

func checkParallel(model Model, history [][]entry, computeInfo bool, timeout time.Duration) (CheckResult, LinearizationInfo) {
	_ = "STUB: not implemented"
	return *new(CheckResult), *new(LinearizationInfo)
}

// if we time out, we might get a false positive

// make sure we've waited for all goroutines to finish,
// otherwise we might race on access to longest[]

// return longest linearizable prefixes that include each history element

// turn longest into a set of unique linearizations

func checkEvents(model Model, history []Event, verbose bool, timeout time.Duration) (CheckResult, LinearizationInfo) {
	_ = "STUB: not implemented"
	return *new(CheckResult), *new(LinearizationInfo)
}

func checkOperations(model Model, history []Operation, verbose bool, timeout time.Duration) (CheckResult, LinearizationInfo) {
	_ = "STUB: not implemented"
	return *new(CheckResult), *new(LinearizationInfo)
}
