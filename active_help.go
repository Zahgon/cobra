package cobra

const (
	activeHelpMarker = "_activeHelp_ "

	activeHelpEnvVarSuffix  = "ACTIVE_HELP"
	activeHelpGlobalEnvVar  = configEnvVarGlobalPrefix + "_" + activeHelpEnvVarSuffix
	activeHelpGlobalDisable = "0"
)

func AppendActiveHelp(compArray []Completion, activeHelpStr string) []Completion {
	_ = "STUB: not implemented"
	return nil
}

func GetActiveHelpConfig(cmd *Command) string { _ = "STUB: not implemented"; return "" }

func activeHelpEnvVar(name string) string { _ = "STUB: not implemented"; return "" }
