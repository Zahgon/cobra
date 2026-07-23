package doc

import (
	"io"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GenManTree(cmd *cobra.Command, header *GenManHeader, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

func GenManTreeFromOpts(cmd *cobra.Command, opts GenManTreeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type GenManTreeOptions struct {
	Header           *GenManHeader
	Path             string
	CommandSeparator string
}

type GenManHeader struct {
	Title   string
	Section string
	Date    *time.Time
	date    string
	Source  string
	Manual  string
}

func GenMan(cmd *cobra.Command, header *GenManHeader, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func fillHeader(header *GenManHeader, name string, disableAutoGen bool) error {
	_ = "STUB: not implemented"
	return nil
}

func manPreamble(buf io.StringWriter, header *GenManHeader, cmd *cobra.Command, dashedName string) {
	_ = "STUB: not implemented"
	return
}

func manPrintFlags(buf io.StringWriter, flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func manPrintOptions(buf io.StringWriter, command *cobra.Command) {
	_ = "STUB: not implemented"
	return
}

func genMan(cmd *cobra.Command, header *GenManHeader) []byte { _ = "STUB: not implemented"; return nil }
