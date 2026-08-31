package cobra

import (
	"io"
)

func (c *Command) genBashCompletion(w io.Writer, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genBashComp(buf io.StringWriter, name string, includeDesc bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) GenBashCompletionFileV2(filename string, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenBashCompletionV2(w io.Writer, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}
