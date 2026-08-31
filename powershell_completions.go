package cobra

import (
	"io"
)

func genPowerShellComp(buf io.StringWriter, name string, includeDesc bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) genPowerShellCompletion(w io.Writer, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) genPowerShellCompletionFile(filename string, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenPowerShellCompletionFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenPowerShellCompletion(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c *Command) GenPowerShellCompletionFileWithDesc(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenPowerShellCompletionWithDesc(w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
