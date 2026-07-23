package cobra

import (
	flag "github.com/spf13/pflag"
)

const (
	requiredAsGroupAnnotation   = "cobra_annotation_required_if_others_set"
	oneRequiredAnnotation       = "cobra_annotation_one_required"
	mutuallyExclusiveAnnotation = "cobra_annotation_mutually_exclusive"
)

func (c *Command) MarkFlagsRequiredTogether(flagNames ...string) { _ = "STUB: not implemented"; return }

func (c *Command) MarkFlagsOneRequired(flagNames ...string) { _ = "STUB: not implemented"; return }

func (c *Command) MarkFlagsMutuallyExclusive(flagNames ...string) {
	_ = "STUB: not implemented"
	return
}

func (c *Command) ValidateFlagGroups() error { _ = "STUB: not implemented"; return nil }

func hasAllFlags(fs *flag.FlagSet, flagnames ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func processFlagForGroupAnnotation(flags *flag.FlagSet, pflag *flag.Flag, annotation string, groupStatus map[string]map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func validateRequiredFlagGroups(data map[string]map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOneRequiredFlagGroups(data map[string]map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExclusiveFlagGroups(data map[string]map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func sortedKeys(m map[string]map[string]bool) []string { _ = "STUB: not implemented"; return nil }

func (c *Command) enforceFlagGroupsForCompletion() { _ = "STUB: not implemented"; return }
