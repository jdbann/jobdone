package teast

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// CmdsEqual checks whether the provided tea.Cmd functions are both nil or both
// return the same value.
func CmdsEqual(a, b tea.Cmd) bool {
	// If both are nil, commands are equal
	if a == nil && b == nil {
		return true
	}

	// If either is nil, commands are not equal
	if a == nil || b == nil {
		return false
	}

	// Get the messages returned from the commands
	aMsg, bMsg := a(), b()

	// If both are batches, check if the batches are equal
	if isCmdBatch(aMsg) && isCmdBatch(bMsg) {
		return batchCmdsEqual(aMsg, bMsg)
	}

	// Check if the messages match
	return reflect.DeepEqual(aMsg, bMsg)
}

var cmdType = reflect.TypeOf(tea.Cmd(func() tea.Msg { return nil }))
var batchType = reflect.SliceOf(cmdType)

func isCmdBatch(msg tea.Msg) bool {
	msgValue := reflect.ValueOf(msg)
	return msgValue.CanConvert(batchType)
}

func batchCmdsEqual(a, b tea.Msg) bool {
	// Get the underlying values of the messages
	aValue, bValue := reflect.ValueOf(a), reflect.ValueOf(b)

	// Convert the values to tea.Cmd slices
	aBatch := aValue.Convert(batchType).Interface().([]tea.Cmd)
	bBatch := bValue.Convert(batchType).Interface().([]tea.Cmd)

	// Not equal if they have different lengths - tea.Batch will strip nils
	if len(aBatch) != len(bBatch) {
		return false
	}

	// Check that the commands in the batches are equal
	for i := range aBatch {
		if !CmdsEqual(aBatch[i], bBatch[i]) {
			return false
		}
	}

	return true
}
