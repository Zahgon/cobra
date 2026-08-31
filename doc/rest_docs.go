package doc

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
)

func printOptionsReST(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultLinkHandler(name, ref string) string { _ = "STUB: not implemented"; return "" }

func GenReST(cmd *cobra.Command, w io.Writer) error { _ = "STUB: not implemented"; return nil }

func GenReSTCustom(cmd *cobra.Command, w io.Writer, linkHandler func(string, string) string) error {
	_ = "STUB: not implemented"
	return nil
}

func GenReSTTree(cmd *cobra.Command, dir string) error { _ = "STUB: not implemented"; return nil }

func GenReSTTreeCustom(cmd *cobra.Command, dir string, filePrepender func(string) string, linkHandler func(string, string) string) error {
	_ = "STUB: not implemented"
	return nil
}

func indentString(s, p string) string { _ = "STUB: not implemented"; return "" }
