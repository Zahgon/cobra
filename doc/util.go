package doc

import (
	"github.com/spf13/cobra"
)

func hasSeeAlso(cmd *cobra.Command) bool { _ = "STUB: not implemented"; return false }

func forceMultiLine(s string) string { _ = "STUB: not implemented"; return "" }

type byName []*cobra.Command

func (s byName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byName) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
