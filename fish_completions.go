package cobra

import (
	"io"
)

func genFishComp(buf io.StringWriter, name string, includeDesc bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) GenFishCompletion(w io.Writer, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenFishCompletionFile(filename string, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}
