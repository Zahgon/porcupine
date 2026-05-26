package porcupine

import (
	"embed"
	"io"
)

type historyElement struct {
	ClientId      int
	Start         int
	OriginalStart string
	End           int
	OriginalEnd   string
	Description   string
	Metadata      string
}

type annotation struct {
	ClientId        int
	Tag             string
	Start           int
	End             int
	Description     string
	Details         string
	Annotation      bool // always true
	TextColor       string
	BackgroundColor string
}

type linearizationStep struct {
	Index            int
	StateDescription string
}

type partialLinearization = []linearizationStep

type partitionVisualizationData struct {
	History               []historyElement
	PartialLinearizations []partialLinearization
	Largest               map[int]int
}

type visualizationData struct {
	Partitions  []partitionVisualizationData
	Annotations []annotation
}

// Annotations to add to histories.
//
// Either a ClientId or Tag must be supplied. The End is optional, for "point
// in time" annotations. If the end is left unspecified, the framework
// interprets it as Start. The text given in Description is shown in the main
// visualization, and the text given in Details (optional) is shown in the
// tooltip for the annotation. TextColor and BackgroundColor are both optional;
// if specified, they should be valid CSS colors, e.g., "#efaefc".
//
// To attach annotations to a visualization, use
// [LinearizationInfo.AddAnnotations].
type Annotation struct {
	ClientId        int
	Tag             string
	Start           int64
	End             int64
	Description     string
	Details         string
	TextColor       string
	BackgroundColor string
	_               struct{} // disallow positional literals, for extensibility
}

// AddAnnotations adds extra annotations to a visualization.
//
// This can be used to add extra client operations  or it can be used to add
// standalone annotations with arbitrary tags, e.g., associated with "servers"
// rather than clients, or even a "test framework".
//
// See documentation on [Annotation] for what kind of annotations you can add.
func (li *LinearizationInfo) AddAnnotations(annotations []Annotation) {
	_ = "STUB: not implemented"
	return
}

// timestampMapping applies a monotonic map to compress timestamps.
//
// This function applies a monotonic map to timestamps so that the encoding of
// timestamps in JSON keeps integers smaller than Number.MAX_SAFE_INTEGER.
// Additionally, this function ensures that the minimum delta between any two
// timestamps is at least 100, to coordinate with index.js, where it is
// convenient to be able to adjust timestamps by an epsilon value (epsilon = 16)
// without them overlapping with other adjusted timestamps.
func timestampMapping(info LinearizationInfo) map[int64]int {
	_ = "STUB: not implemented"
	// find all timestamps
	return nil
}

// sort

// construct mapping

// ensure minimum delta of 100 between timestamps

func computeVisualizationData(model Model, info LinearizationInfo) visualizationData {
	_ = "STUB: not implemented"
	return *new(visualizationData)
}

// history

// prefer return metadata over call metadata

// historyElement.Annotation defaults to false, so we
// don't need to explicitly set it here; all of these
// are non-annotation elements

// partial linearizations

// Visualize produces a visualization of a history and (partial) linearization
// as an HTML file that can be viewed in a web browser.
//
// If the history is linearizable, the visualization shows the linearization of
// the history. If the history is not linearizable, the visualization shows
// partial linearizations and illegal linearization points.
//
// To get the LinearizationInfo that this function requires, you can use
// [CheckOperationsVerbose] / [CheckEventsVerbose].
//
// This function writes the visualization, an HTML file with embedded
// JavaScript and data, to the given output.
func Visualize(model Model, info LinearizationInfo, output io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// VisualizePath is a wrapper around [Visualize] to write the visualization to
// a file path.
func VisualizePath(model Model, info LinearizationInfo, path string) error {
	_ = "STUB: not implemented"
	return nil
}

//go:embed visualization
var visualizationFS embed.FS
