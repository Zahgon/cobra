package cobra

import (
	"github.com/spf13/pflag"
)

func (c *Command) MarkFlagRequired(name string) error { _ = "STUB: not implemented"; return nil }

func (c *Command) MarkPersistentFlagRequired(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func MarkFlagRequired(flags *pflag.FlagSet, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) MarkFlagFilename(name string, extensions ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) MarkFlagCustom(name string, f string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) MarkPersistentFlagFilename(name string, extensions ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func MarkFlagFilename(flags *pflag.FlagSet, name string, extensions ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func MarkFlagCustom(flags *pflag.FlagSet, name string, f string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) MarkFlagDirname(name string) error { _ = "STUB: not implemented"; return nil }

func (c *Command) MarkPersistentFlagDirname(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func MarkFlagDirname(flags *pflag.FlagSet, name string) error {
	_ = "STUB: not implemented"
	return nil
}
