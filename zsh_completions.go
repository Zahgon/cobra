package cobra

import (
	"io"
)

func (c *Command) GenZshCompletionFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenZshCompletion(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c *Command) GenZshCompletionFileNoDesc(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GenZshCompletionNoDesc(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c *Command) MarkZshCompPositionalArgumentFile(argPosition int, patterns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) MarkZshCompPositionalArgumentWords(argPosition int, words ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) genZshCompletionFile(filename string, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) genZshCompletion(w io.Writer, includeDesc bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genZshComp(buf io.StringWriter, name string, includeDesc bool) {
	_ = "STUB: not implemented"
	return
}
