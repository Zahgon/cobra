package cobra

import (
	"regexp"
	"sync"

	"github.com/spf13/pflag"
)

const (
	ShellCompRequestCmd = "__complete"

	ShellCompNoDescRequestCmd = "__completeNoDesc"
)

var flagCompletionFunctions = map[*pflag.Flag]CompletionFunc{}

var flagCompletionMutex = &sync.RWMutex{}

type ShellCompDirective int

type flagCompError struct {
	subCommand string
	flagName   string
}

func (e *flagCompError) Error() string { _ = "STUB: not implemented"; return "" }

const (
	ShellCompDirectiveError ShellCompDirective = 1 << iota

	ShellCompDirectiveNoSpace

	ShellCompDirectiveNoFileComp

	ShellCompDirectiveFilterFileExt

	ShellCompDirectiveFilterDirs

	ShellCompDirectiveKeepOrder

	shellCompDirectiveMaxValue

	ShellCompDirectiveDefault ShellCompDirective = 0
)

const (
	compCmdName              = "completion"
	compCmdNoDescFlagName    = "no-descriptions"
	compCmdNoDescFlagDesc    = "disable completion descriptions"
	compCmdNoDescFlagDefault = false
)

type CompletionOptions struct {
	DisableDefaultCmd bool

	DisableNoDescFlag bool

	DisableDescriptions bool

	HiddenDefaultCmd bool

	DefaultShellCompDirective *ShellCompDirective
}

func (receiver *CompletionOptions) SetDefaultShellCompDirective(directive ShellCompDirective) {
	_ = "STUB: not implemented"
	return
}

type Completion = string

type CompletionFunc = func(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective)

func CompletionWithDesc(choice string, description string) Completion {
	_ = "STUB: not implemented"
	return *new(Completion)
}

func NoFileCompletions(cmd *Command, args []string, toComplete string) ([]Completion, ShellCompDirective) {
	_ = "STUB: not implemented"
	return nil, *new(ShellCompDirective)
}

func FixedCompletions(choices []Completion, directive ShellCompDirective) CompletionFunc {
	_ = "STUB: not implemented"
	return *new(CompletionFunc)
}

func (c *Command) RegisterFlagCompletionFunc(flagName string, f CompletionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) GetFlagCompletionFunc(flagName string) (CompletionFunc, bool) {
	_ = "STUB: not implemented"
	return *new(CompletionFunc), false
}

func (d ShellCompDirective) string() string { _ = "STUB: not implemented"; return "" }

func (c *Command) initCompleteCmd(args []string) { _ = "STUB: not implemented"; return }

type SliceValue interface {
	GetSlice() []string
}

func (c *Command) getCompletions(args []string) (*Command, []Completion, ShellCompDirective, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(ShellCompDirective), nil
}

func helpOrVersionFlagPresent(cmd *Command) bool { _ = "STUB: not implemented"; return false }

func getFlagNameCompletions(flag *pflag.Flag, toComplete string) []Completion {
	_ = "STUB: not implemented"
	return nil
}

func completeRequireFlags(finalCmd *Command, toComplete string) []Completion {
	_ = "STUB: not implemented"
	return nil
}

func checkIfFlagCompletion(finalCmd *Command, args []string, lastArg string) (*pflag.Flag, []string, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func (c *Command) InitDefaultCompletionCmd(args ...string) { _ = "STUB: not implemented"; return }

func findFlag(cmd *Command, name string) *pflag.Flag { _ = "STUB: not implemented"; return nil }

func CompDebug(msg string, printToStdErr bool) { _ = "STUB: not implemented"; return }

//nolint:gosec // G703:BASH_COMP_DEBUG_FILE intentionally user-controlled for completion debug logging.

func CompDebugln(msg string, printToStdErr bool) { _ = "STUB: not implemented"; return }

func CompError(msg string) { _ = "STUB: not implemented"; return }

func CompErrorln(msg string) { _ = "STUB: not implemented"; return }

const (
	configEnvVarGlobalPrefix       = "COBRA"
	configEnvVarSuffixDescriptions = "COMPLETION_DESCRIPTIONS"
)

var configEnvVarPrefixSubstRegexp = regexp.MustCompile(`[^A-Z0-9_]`)

func configEnvVar(name, suffix string) string { _ = "STUB: not implemented"; return "" }

func getEnvConfig(cmd *Command, suffix string) string { _ = "STUB: not implemented"; return "" }
