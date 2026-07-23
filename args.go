package cobra

type PositionalArgs func(cmd *Command, args []string) error

func legacyArgs(cmd *Command, args []string) error { _ = "STUB: not implemented"; return nil }

func NoArgs(cmd *Command, args []string) error { _ = "STUB: not implemented"; return nil }

func OnlyValidArgs(cmd *Command, args []string) error { _ = "STUB: not implemented"; return nil }

func NoDuplicateArgs(cmd *Command, args []string) error { _ = "STUB: not implemented"; return nil }

func ArbitraryArgs(cmd *Command, args []string) error { _ = "STUB: not implemented"; return nil }

func MinimumNArgs(n int) PositionalArgs { _ = "STUB: not implemented"; return *new(PositionalArgs) }

func MaximumNArgs(n int) PositionalArgs { _ = "STUB: not implemented"; return *new(PositionalArgs) }

func ExactArgs(n int) PositionalArgs { _ = "STUB: not implemented"; return *new(PositionalArgs) }

func RangeArgs(min int, max int) PositionalArgs {
	_ = "STUB: not implemented"
	return *new(PositionalArgs)
}

func MatchAll(pargs ...PositionalArgs) PositionalArgs {
	_ = "STUB: not implemented"
	return *new(PositionalArgs)
}

func ExactValidArgs(n int) PositionalArgs { _ = "STUB: not implemented"; return *new(PositionalArgs) }
