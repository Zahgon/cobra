package cobra

import (
	"io"
	"strings"
	"text/template"
	"time"
)

var templateFuncs = template.FuncMap{
	"trim":                    strings.TrimSpace,
	"trimRightSpace":          trimRightSpace,
	"trimTrailingWhitespaces": trimRightSpace,
	"appendIfNotPresent":      appendIfNotPresent,
	"rpad":                    rpad,
	"gt":                      Gt,
	"eq":                      Eq,
}

var initializers []func()
var finalizers []func()

const (
	defaultPrefixMatching   = false
	defaultCommandSorting   = true
	defaultCaseInsensitive  = false
	defaultTraverseRunHooks = false
)

var EnablePrefixMatching = defaultPrefixMatching

var EnableCommandSorting = defaultCommandSorting

var EnableCaseInsensitive = defaultCaseInsensitive

var EnableTraverseRunHooks = defaultTraverseRunHooks

var MousetrapHelpText = `This is a command line tool.

You need to open cmd.exe and run it from there.
`

var MousetrapDisplayDuration = 5 * time.Second

func AddTemplateFunc(name string, tmplFunc interface{}) { _ = "STUB: not implemented"; return }

func AddTemplateFuncs(tmplFuncs template.FuncMap) { _ = "STUB: not implemented"; return }

func OnInitialize(y ...func()) { _ = "STUB: not implemented"; return }

func OnFinalize(y ...func()) { _ = "STUB: not implemented"; return }

func Gt(a interface{}, b interface{}) bool { _ = "STUB: not implemented"; return false }

func Eq(a interface{}, b interface{}) bool { _ = "STUB: not implemented"; return false }

func trimRightSpace(s string) string { _ = "STUB: not implemented"; return "" }

func appendIfNotPresent(s, stringToAppend string) string { _ = "STUB: not implemented"; return "" }

func rpad(s string, padding int) string { _ = "STUB: not implemented"; return "" }

func tmpl(text string) *tmplFunc { _ = "STUB: not implemented"; return nil }

func ld(s, t string, ignoreCase bool) int { _ = "STUB: not implemented"; return 0 }

func stringInSlice(a string, list []string) bool { _ = "STUB: not implemented"; return false }

func CheckErr(msg interface{}) { _ = "STUB: not implemented"; return }

func WriteStringAndCheck(b io.StringWriter, s string) { _ = "STUB: not implemented"; return }
