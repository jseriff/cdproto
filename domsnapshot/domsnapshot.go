// Package domsnapshot provides the Chrome Debugging Protocol
// commands, types, and events for the DOMSnapshot domain.
//
// This domain facilitates obtaining document snapshots with DOM, layout, and
// style information.
//
// Generated by the cdproto-gen command.
package domsnapshot

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// DisableParams disables DOM snapshot agent for the given page.
type DisableParams struct{}

// Disable disables DOM snapshot agent for the given page.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes DOMSnapshot.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables DOM snapshot agent for the given page.
type EnableParams struct{}

// Enable enables DOM snapshot agent for the given page.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes DOMSnapshot.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, nil, nil)
}

// CaptureSnapshotParams returns a document snapshot, including the full DOM
// tree of the root node (including iframes, template contents, and imported
// documents) in a flattened array, as well as layout and white-listed computed
// style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
type CaptureSnapshotParams struct {
	ComputedStyles []string `json:"computedStyles"` // Whitelist of computed styles to return.
}

// CaptureSnapshot returns a document snapshot, including the full DOM tree
// of the root node (including iframes, template contents, and imported
// documents) in a flattened array, as well as layout and white-listed computed
// style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
//
// parameters:
//   computedStyles - Whitelist of computed styles to return.
func CaptureSnapshot(computedStyles []string) *CaptureSnapshotParams {
	return &CaptureSnapshotParams{
		ComputedStyles: computedStyles,
	}
}

// CaptureSnapshotReturns return values.
type CaptureSnapshotReturns struct {
	Nodes   *DOMTreeSnapshot    `json:"nodes,omitempty"`   // The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Layout  *LayoutTreeSnapshot `json:"layout,omitempty"`  // The nodes in the layout tree.
	Strings []string            `json:"strings,omitempty"` // Shared string table that all string properties refer to with indexes.
}

// Do executes DOMSnapshot.captureSnapshot against the provided context.
//
// returns:
//   nodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
//   layout - The nodes in the layout tree.
//   strings - Shared string table that all string properties refer to with indexes.
func (p *CaptureSnapshotParams) Do(ctxt context.Context, h cdp.Executor) (nodes *DOMTreeSnapshot, layout *LayoutTreeSnapshot, strings []string, err error) {
	// execute
	var res CaptureSnapshotReturns
	err = h.Execute(ctxt, CommandCaptureSnapshot, p, &res)
	if err != nil {
		return nil, nil, nil, err
	}

	return res.Nodes, res.Layout, res.Strings, nil
}

// Command names.
const (
	CommandDisable         = "DOMSnapshot.disable"
	CommandEnable          = "DOMSnapshot.enable"
	CommandCaptureSnapshot = "DOMSnapshot.captureSnapshot"
)
