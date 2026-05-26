package porcupine

// An Operation is an element of a history.
//
// This package supports two different representations of histories, as a
// sequence of Operation or [Event]. In the Operation representation, function
// call/returns are packaged together, along with timestamps of when the
// function call was made and when the function call returned.
//
// The interval [Call, Return] is interpreted as a closed interval, so an
// operation with interval [10, 20] is concurrent with another operation with
// interval [20, 30].
type Operation struct {
	ClientId int // optional, unless you want a visualization; zero-indexed
	Input    interface{}
	Call     int64 // invocation timestamp
	Output   interface{}
	Return   int64 // response timestamp
	// Metadata contains arbitrary metadata associated with the operation.
	// It is not used for linearizability checking but can be used for visualization.
	Metadata interface{}
	_        struct{} // disallow positional literals, for extensibility
}

// Interpreting the interval [Call, Return] as a closed interval is the only
// reasonable approach for how we expect this library to be used. Otherwise, we
// might have the following situation, when using a monotonic clock to get
// timestamps (where successive calls to the clock return values that are always
// greater than _or equal to_ previously returned values):
//
// - Client 1 calls clock(), gets ts=1, invokes put("x", "y")
// - Client 2 calls clock(), gets ts=2, invokes get("x")
// - Client 1 operation returns, calls clock(), gets ts=2
// - Client 2 operation returns "", calls clock(), gets ts=3
//
// These operations were concurrent, but if we interpret the intervals as
// half-open, for example, Client 1's operation had interval [1, 2) and Client
// 2's operation had interval [2, 3), so they are not concurrent operations, and
// we'd say that this history is not linearizable, which is not correct. The
// only sensible approach is to interpret the interval [Call, Return] as a
// closed interval.

// An EventKind tags an [Event] as either a function call or a return.
type EventKind bool

const (
	CallEvent   EventKind = false
	ReturnEvent EventKind = true
)

// An Event is an element of a history, a function call event or a return
// event.
//
// This package supports two different representations of histories, as a
// sequence of Event or [Operation]. In the Event representation, function
// calls and returns are only relatively ordered and do not have absolute
// timestamps.
//
// The Id field is used to match a function call event with its corresponding
// return event.
type Event struct {
	ClientId int // optional, unless you want a visualization; zero-indexed
	Kind     EventKind
	Value    interface{}
	Id       int
	// Metadata contains arbitrary metadata associated with the operation.
	// It is not used for linearizability checking but can be used for visualization.
	// Can be set on CallEvent or ReturnEvent. If both have metadata, ReturnEvent metadata takes precedence.
	Metadata interface{}
	_        struct{} // disallow positional literals, for extensibility
}

// A Model is a sequential specification of a system.
//
// Note: models in this package are expected to be purely functional. That is,
// the model Step function should not modify the given state (or input or
// output), but return a new state.
//
// Only the Init, Step, and Equal functions are necessary to specify if you
// just want to test histories for linearizability.
//
// Implementing the partition functions can greatly improve performance. If
// you're implementing the partition function, the model Init and Step
// functions can be per-partition. For example, if your specification is for a
// key-value store and you partition by key, then the per-partition state
// representation can just be a single value rather than a map.
//
// Implementing DescribeOperation and DescribeState will produce nicer
// visualizations.
//
// It may be helpful to look at this package's [test code] for examples of how
// to write models, including models that include partition functions.
//
// [test code]: https://github.com/anishathalye/porcupine/blob/master/porcupine_test.go
type Model struct {
	// Partition functions, such that a history is linearizable if and only
	// if each partition is linearizable. If left nil, this package will
	// skip partitioning.
	Partition      func(history []Operation) [][]Operation
	PartitionEvent func(history []Event) [][]Event
	// Initial state of the system.
	Init func() interface{}
	// Step function for the system. Returns whether or not the system
	// could take this step with the given inputs and outputs and also
	// returns the new state. This function must be a pure function: it
	// cannot mutate the given state.
	Step func(state interface{}, input interface{}, output interface{}) (bool, interface{})
	// Equality on states. If left nil, this package will use == as a
	// fallback ([ShallowEqual]).
	Equal func(state1, state2 interface{}) bool
	// For visualization, describe an operation as a string. For example,
	// "Get('x') -> 'y'". Can be omitted if you're not producing
	// visualizations.
	DescribeOperation func(input interface{}, output interface{}) string
	// For visualization purposes, describe a state as a string. For
	// example, "{'x' -> 'y', 'z' -> 'w'}". Can be omitted if you're not
	// producing visualizations.
	DescribeState func(state interface{}) string
	// For visualization purposes, describe metadata as a string. Can be
	// omitted if you're not producing visualizations.
	DescribeOperationMetadata func(info interface{}) string
	_                         struct{} // disallow positional literals, for extensibility
}

// A NondeterministicModel is a nondeterministic sequential specification of a
// system.
//
// For basics on models, see the documentation for [Model].  In contrast to
// Model, NondeterministicModel has a step function that returns a set of
// states, indicating all possible next states. It can be converted to a Model
// using the [NondeterministicModel.ToModel] function.
//
// It may be helpful to look at this package's [test code] for examples of how
// to write and use nondeterministic models.
//
// [test code]: https://github.com/anishathalye/porcupine/blob/master/porcupine_test.go
type NondeterministicModel struct {
	// Partition functions, such that a history is linearizable if and only
	// if each partition is linearizable. If left nil, this package will
	// skip partitioning.
	Partition      func(history []Operation) [][]Operation
	PartitionEvent func(history []Event) [][]Event
	// Initial states of the system.
	Init func() []interface{}
	// Step function for the system. Returns all possible next states for
	// the given state, input, and output. If the system cannot step with
	// the given state/input to produce the given output, this function
	// should return an empty slice.
	Step func(state interface{}, input interface{}, output interface{}) []interface{}
	// Equality on states. If left nil, this package will use == as a
	// fallback ([ShallowEqual]).
	Equal func(state1, state2 interface{}) bool
	// For visualization, describe an operation as a string. For example,
	// "Get('x') -> 'y'". Can be omitted if you're not producing
	// visualizations.
	DescribeOperation func(input interface{}, output interface{}) string
	// For visualization purposes, describe a state as a string. For
	// example, "{'x' -> 'y', 'z' -> 'w'}". Can be omitted if you're not
	// producing visualizations.
	DescribeState func(state interface{}) string
	// For visualization purposes, describe metadata as a string. Can be
	// omitted if you're not producing visualizations.
	DescribeOperationMetadata func(info interface{}) string
	_                         struct{} // disallow positional literals, for extensibility
}

func merge(states []interface{}, eq func(state1, state2 interface{}) bool) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ToModel converts a [NondeterministicModel] to a [Model] using a power set
// construction.
//
// This makes it suitable for use in linearizability checking operations like
// [CheckOperations]. This is a general construction that can be used for any
// nondeterministic model. It relies on the NondeterministicModel's Equal
// function to merge states. You may be able to achieve better performance by
// implementing a Model directly.
func (nm *NondeterministicModel) ToModel() Model {
	_ = "STUB: not implemented"
	// like fillDefault
	return *new(Model)
}

// we need this wrapper to convert a []interface{} to an interface{}

// this operates on sets of states that have been merged, so we
// don't need to check inclusion in both directions

// noPartition is a fallback partition function that partitions the history
// into a single partition containing all of the operations.
func noPartition(history []Operation) [][]Operation { _ = "STUB: not implemented"; return nil }

// noPartitionEvent is a fallback partition function that partitions the
// history into a single partition containing all of the events.
func noPartitionEvent(history []Event) [][]Event { _ = "STUB: not implemented"; return nil }

// shallowEqual is a fallback equality function that compares two states using
// ==.
func shallowEqual(state1, state2 interface{}) bool { _ = "STUB: not implemented"; return false }

// defaultDescribeOperation is a fallback to convert an operation to a string.
// It renders inputs and outputs using the "%v" format specifier.
func defaultDescribeOperation(input interface{}, output interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultDescribeState is a fallback to convert a state to a string. It
// renders the state using the "%v" format specifier.
func defaultDescribeState(state interface{}) string { _ = "STUB: not implemented"; return "" }

// defaultDescribeOperationMetadata is a fallback to convert metadata to a
// string. It renders the metadata using the "%v" format specifier.
func defaultDescribeOperationMetadata(info interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// A CheckResult is the result of a linearizability check.
//
// Checking for linearizability is decidable, but it is an NP-hard problem, so
// the checker might take a long time. If a timeout is not given, functions in
// this package will always return Ok or Illegal, but if a timeout is supplied,
// then some functions may return Unknown. Depending on the use case, you can
// interpret an Unknown result as Ok (i.e., the tool didn't find a
// linearizability violation within the given timeout).
type CheckResult string

const (
	Unknown CheckResult = "Unknown" // timed out
	Ok      CheckResult = "Ok"
	Illegal CheckResult = "Illegal"
)
