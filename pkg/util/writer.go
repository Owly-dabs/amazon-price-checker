package util

import (
	"os"
	"text/tabwriter"
)

// DEFAULTS
var (
	minWidth      = 0
	tabWidth      = 8
	padding       = 1
	padChar  byte = '\t'
	flags         = 0
)

type TabWriter struct {
	writer *tabwriter.Writer
}

// NewTabWriter creates a new TabWriter with the specified output, minwidth, tabwidth,
// padding, padchar, and flags.
func NewTabWriter(minwidth, tabwidth, padding int, padchar byte, flags uint) *TabWriter {
	return &TabWriter{
		writer: tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags),
	}
}

// Write writes the given data to the underlying tabwriter.Writer.
func (tw *TabWriter) Write(data []byte) (int, error) {
	return tw.writer.Write(data)
}

// Flush flushes the underlying tabwriter.Writer.
func (tw *TabWriter) Flush() error {
	return tw.writer.Flush()
}
