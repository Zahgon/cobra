//go:build windows
// +build windows

package cobra

var preExecHookFn = preExecHook

func preExecHook(c *Command) { _ = "STUB: not implemented"; return }
