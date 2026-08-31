package cobra

import (
	"io"

	"github.com/spf13/pflag"
)

const (
	BashCompFilenameExt     = "cobra_annotation_bash_completion_filename_extensions"
	BashCompCustom          = "cobra_annotation_bash_completion_custom"
	BashCompOneRequiredFlag = "cobra_annotation_bash_completion_one_required_flag"
	BashCompSubdirsInDir    = "cobra_annotation_bash_completion_subdirs_in_dir"
)

func writePreamble(buf io.StringWriter, name string) { _ = "STUB: not implemented"; return }

func writePostscript(buf io.StringWriter, name string) { _ = "STUB: not implemented"; return }

func writeCommands(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func writeFlagHandler(buf io.StringWriter, name string, annotations map[string][]string, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

const cbn = "\")\n"

func writeShortFlag(buf io.StringWriter, flag *pflag.Flag, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

func writeFlag(buf io.StringWriter, flag *pflag.Flag, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

func writeLocalNonPersistentFlag(buf io.StringWriter, flag *pflag.Flag) {
	_ = "STUB: not implemented"
	return
}

func prepareCustomAnnotationsForFlags(cmd *Command) { _ = "STUB: not implemented"; return }

func writeFlags(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func writeRequiredFlag(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func writeRequiredNouns(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func writeCmdAliases(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func writeArgAliases(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func gen(buf io.StringWriter, cmd *Command) { _ = "STUB: not implemented"; return }

func (c *Command) GenBashCompletion(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func nonCompletableFlag(flag *pflag.Flag) bool { _ = "STUB: not implemented"; return false }

func (c *Command) GenBashCompletionFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}
