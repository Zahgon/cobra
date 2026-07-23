package cobra

import (
	"bytes"
	"context"
	"io"

	flag "github.com/spf13/pflag"
)

const (
	FlagSetByCobraAnnotation     = "cobra_annotation_flag_set_by_cobra"
	CommandDisplayNameAnnotation = "cobra_annotation_command_display_name"

	helpFlagName    = "help"
	helpCommandName = "help"
)

type FParseErrWhitelist flag.ParseErrorsAllowlist

type Group struct {
	ID    string
	Title string
}

type Command struct {
	Use string

	Aliases []string

	SuggestFor []string

	Short string

	GroupID string

	Long string

	Example string

	ValidArgs []Completion

	ValidArgsFunction CompletionFunc

	Args PositionalArgs

	ArgAliases []string

	BashCompletionFunction string

	Deprecated string

	Annotations map[string]string

	Version string

	PersistentPreRun func(cmd *Command, args []string)

	PersistentPreRunE func(cmd *Command, args []string) error

	PreRun func(cmd *Command, args []string)

	PreRunE func(cmd *Command, args []string) error

	Run func(cmd *Command, args []string)

	RunE func(cmd *Command, args []string) error

	PostRun func(cmd *Command, args []string)

	PostRunE func(cmd *Command, args []string) error

	PersistentPostRun func(cmd *Command, args []string)

	PersistentPostRunE func(cmd *Command, args []string) error

	commandgroups []*Group

	args []string

	flagErrorBuf *bytes.Buffer

	flags *flag.FlagSet

	pflags *flag.FlagSet

	lflags *flag.FlagSet

	iflags *flag.FlagSet

	parentsPflags *flag.FlagSet

	globNormFunc func(f *flag.FlagSet, name string) flag.NormalizedName

	usageFunc func(*Command) error

	usageTemplate *tmplFunc

	flagErrorFunc func(*Command, error) error

	helpTemplate *tmplFunc

	helpFunc func(*Command, []string)

	helpCommand *Command

	helpCommandGroupID string

	completionCommandGroupID string

	versionTemplate *tmplFunc

	errPrefix string

	inReader io.Reader

	outWriter io.Writer

	errWriter io.Writer

	FParseErrWhitelist FParseErrWhitelist

	CompletionOptions CompletionOptions

	commandsAreSorted bool

	commandCalledAs struct {
		name   string
		called bool
	}

	ctx context.Context

	commands []*Command

	parent *Command

	commandsMaxUseLen         int
	commandsMaxCommandPathLen int
	commandsMaxNameLen        int

	TraverseChildren bool

	Hidden bool

	SilenceErrors bool

	SilenceUsage bool

	DisableFlagParsing bool

	DisableAutoGenTag bool

	DisableFlagsInUseLine bool

	DisableSuggestions bool

	SuggestionsMinimumDistance int
}

func (c *Command) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *Command) SetContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Command) SetArgs(a []string) { _ = "STUB: not implemented"; return }

func (c *Command) SetOutput(output io.Writer) { _ = "STUB: not implemented"; return }

func (c *Command) SetOut(newOut io.Writer) { _ = "STUB: not implemented"; return }

func (c *Command) SetErr(newErr io.Writer) { _ = "STUB: not implemented"; return }

func (c *Command) SetIn(newIn io.Reader) { _ = "STUB: not implemented"; return }

func (c *Command) SetUsageFunc(f func(*Command) error) { _ = "STUB: not implemented"; return }

func (c *Command) SetUsageTemplate(s string) { _ = "STUB: not implemented"; return }

func (c *Command) SetFlagErrorFunc(f func(*Command, error) error) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) SetHelpFunc(f func(*Command, []string)) { _ = "STUB: not implemented"; return }

func (c *Command) SetHelpCommand(cmd *Command) { _ = "STUB: not implemented"; return }

func (c *Command) SetHelpCommandGroupID(groupID string) { _ = "STUB: not implemented"; return }

func (c *Command) SetCompletionCommandGroupID(groupID string) { _ = "STUB: not implemented"; return }

func (c *Command) SetHelpTemplate(s string) { _ = "STUB: not implemented"; return }

func (c *Command) SetVersionTemplate(s string) { _ = "STUB: not implemented"; return }

func (c *Command) SetErrPrefix(s string) { _ = "STUB: not implemented"; return }

func (c *Command) SetGlobalNormalizationFunc(n func(f *flag.FlagSet, name string) flag.NormalizedName) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) OutOrStdout() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (c *Command) OutOrStderr() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (c *Command) ErrOrStderr() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (c *Command) InOrStdin() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (c *Command) getOut(def io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (c *Command) getErr(def io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (c *Command) getIn(def io.Reader) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (c *Command) UsageFunc() (f func(*Command) error) { _ = "STUB: not implemented"; return nil }

func (c *Command) getUsageTemplateFunc() func(w io.Writer, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) Usage() error { _ = "STUB: not implemented"; return nil }

func (c *Command) HelpFunc() func(*Command, []string) { _ = "STUB: not implemented"; return nil }

func (c *Command) getHelpTemplateFunc() func(w io.Writer, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) Help() error { _ = "STUB: not implemented"; return nil }

func (c *Command) UsageString() string { _ = "STUB: not implemented"; return "" }

func (c *Command) FlagErrorFunc() (f func(*Command, error) error) {
	_ = "STUB: not implemented"
	return nil
}

const minUsagePadding = 25

func (c *Command) UsagePadding() int { _ = "STUB: not implemented"; return 0 }

const minCommandPathPadding = 11

func (c *Command) CommandPathPadding() int { _ = "STUB: not implemented"; return 0 }

const minNamePadding = 11

func (c *Command) NamePadding() int { _ = "STUB: not implemented"; return 0 }

func (c *Command) UsageTemplate() string { _ = "STUB: not implemented"; return "" }

func (c *Command) HelpTemplate() string { _ = "STUB: not implemented"; return "" }

func (c *Command) VersionTemplate() string { _ = "STUB: not implemented"; return "" }

func (c *Command) getVersionTemplateFunc() func(w io.Writer, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) ErrPrefix() string { _ = "STUB: not implemented"; return "" }

func hasNoOptDefVal(name string, fs *flag.FlagSet) bool { _ = "STUB: not implemented"; return false }

func shortHasNoOptDefVal(name string, fs *flag.FlagSet) bool {
	_ = "STUB: not implemented"
	return false
}

func stripFlags(args []string, c *Command) []string { _ = "STUB: not implemented"; return nil }

func (c *Command) argsMinusFirstX(args []string, x string) []string {
	_ = "STUB: not implemented"
	return nil
}

func isFlagArg(arg string) bool { _ = "STUB: not implemented"; return false }

func (c *Command) Find(args []string) (*Command, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Command) findSuggestions(arg string) string { _ = "STUB: not implemented"; return "" }

func (c *Command) findNext(next string) *Command { _ = "STUB: not implemented"; return nil }

func (c *Command) Traverse(args []string) (*Command, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Command) SuggestionsFor(typedName string) []string { _ = "STUB: not implemented"; return nil }

func (c *Command) VisitParents(fn func(*Command)) { _ = "STUB: not implemented"; return }

func (c *Command) Root() *Command { _ = "STUB: not implemented"; return nil }

func (c *Command) ArgsLenAtDash() int { _ = "STUB: not implemented"; return 0 }

func (c *Command) execute(a []string) (err error) { _ = "STUB: not implemented"; return nil }

func (c *Command) preRun() { _ = "STUB: not implemented"; return }

func (c *Command) postRun() { _ = "STUB: not implemented"; return }

func (c *Command) ExecuteContext(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Command) Execute() error { _ = "STUB: not implemented"; return nil }

func (c *Command) ExecuteContextC(ctx context.Context) (*Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Command) ExecuteC() (cmd *Command, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Command) ValidateArgs(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *Command) ValidateRequiredFlags() error { _ = "STUB: not implemented"; return nil }

func (c *Command) checkCommandGroups() { _ = "STUB: not implemented"; return }

func (c *Command) InitDefaultHelpFlag() { _ = "STUB: not implemented"; return }

func (c *Command) InitDefaultVersionFlag() { _ = "STUB: not implemented"; return }

func (c *Command) InitDefaultHelpCmd() { _ = "STUB: not implemented"; return }

func (c *Command) ResetCommands() { _ = "STUB: not implemented"; return }

type commandSorterByName []*Command

func (c commandSorterByName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (c commandSorterByName) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (c commandSorterByName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c *Command) Commands() []*Command { _ = "STUB: not implemented"; return nil }

func (c *Command) AddCommand(cmds ...*Command) { _ = "STUB: not implemented"; return }

func (c *Command) Groups() []*Group { _ = "STUB: not implemented"; return nil }

func (c *Command) AllChildCommandsHaveGroup() bool { _ = "STUB: not implemented"; return false }

func (c *Command) ContainsGroup(groupID string) bool { _ = "STUB: not implemented"; return false }

func (c *Command) AddGroup(groups ...*Group) { _ = "STUB: not implemented"; return }

func (c *Command) RemoveCommand(cmds ...*Command) { _ = "STUB: not implemented"; return }

func (c *Command) Print(i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) Println(i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) Printf(format string, i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) PrintErr(i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) PrintErrln(i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) PrintErrf(format string, i ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Command) CommandPath() string { _ = "STUB: not implemented"; return "" }

func (c *Command) DisplayName() string { _ = "STUB: not implemented"; return "" }

func (c *Command) UseLine() string { _ = "STUB: not implemented"; return "" }

func (c *Command) DebugFlags() { _ = "STUB: not implemented"; return }

func (c *Command) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Command) HasAlias(s string) bool { _ = "STUB: not implemented"; return false }

func (c *Command) CalledAs() string { _ = "STUB: not implemented"; return "" }

func (c *Command) hasNameOrAliasPrefix(prefix string) bool { _ = "STUB: not implemented"; return false }

func (c *Command) NameAndAliases() string { _ = "STUB: not implemented"; return "" }

func (c *Command) HasExample() bool { _ = "STUB: not implemented"; return false }

func (c *Command) Runnable() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasSubCommands() bool { _ = "STUB: not implemented"; return false }

func (c *Command) IsAvailableCommand() bool { _ = "STUB: not implemented"; return false }

func (c *Command) IsAdditionalHelpTopicCommand() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasHelpSubCommands() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasAvailableSubCommands() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasParent() bool { _ = "STUB: not implemented"; return false }

func (c *Command) GlobalNormalizationFunc() func(f *flag.FlagSet, name string) flag.NormalizedName {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) Flags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) LocalNonPersistentFlags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) LocalFlags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) InheritedFlags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) NonInheritedFlags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) PersistentFlags() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func (c *Command) ResetFlags() { _ = "STUB: not implemented"; return }

func (c *Command) HasFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasPersistentFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasLocalFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasInheritedFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasAvailableFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasAvailablePersistentFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasAvailableLocalFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) HasAvailableInheritedFlags() bool { _ = "STUB: not implemented"; return false }

func (c *Command) Flag(name string) (flag *flag.Flag) { _ = "STUB: not implemented"; return nil }

func (c *Command) persistentFlag(name string) (flag *flag.Flag) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Command) ParseFlags(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *Command) Parent() *Command { _ = "STUB: not implemented"; return nil }

func (c *Command) mergePersistentFlags() { _ = "STUB: not implemented"; return }

func (c *Command) updateParentsPflags() { _ = "STUB: not implemented"; return }

func commandNameMatches(s string, t string) bool { _ = "STUB: not implemented"; return false }

type tmplFunc struct {
	tmpl string
	fn   func(io.Writer, interface{}) error
}

const defaultUsageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

func defaultUsageFunc(w io.Writer, in interface{}) error { _ = "STUB: not implemented"; return nil }

const defaultHelpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

func defaultHelpFunc(w io.Writer, in interface{}) error { _ = "STUB: not implemented"; return nil }

const defaultVersionTemplate = `{{with .DisplayName}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
`

func defaultVersionFunc(w io.Writer, in interface{}) error { _ = "STUB: not implemented"; return nil }
