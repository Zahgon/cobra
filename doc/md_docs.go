package doc

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
)

const markdownExtension = ".md"

func printOptions(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func GenMarkdown(cmd *cobra.Command, w io.Writer) error { _ = "STUB: not implemented"; return nil }

func GenMarkdownCustom(cmd *cobra.Command, w io.Writer, linkHandler func(string) string) error {
	_ = "STUB: not implemented"
	return nil
}

func GenMarkdownTree(cmd *cobra.Command, dir string) error { _ = "STUB: not implemented"; return nil }

func GenMarkdownTreeCustom(cmd *cobra.Command, dir string, filePrepender, linkHandler func(string) string) error {
	_ = "STUB: not implemented"
	return nil
}
