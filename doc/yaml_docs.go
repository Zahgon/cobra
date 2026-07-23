package doc

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type cmdOption struct {
	Name         string
	Shorthand    string `yaml:",omitempty"`
	DefaultValue string `yaml:"default_value,omitempty"`
	Usage        string `yaml:",omitempty"`
}

type cmdDoc struct {
	Name             string
	Synopsis         string      `yaml:",omitempty"`
	Description      string      `yaml:",omitempty"`
	Usage            string      `yaml:",omitempty"`
	Options          []cmdOption `yaml:",omitempty"`
	InheritedOptions []cmdOption `yaml:"inherited_options,omitempty"`
	Example          string      `yaml:",omitempty"`
	SeeAlso          []string    `yaml:"see_also,omitempty"`
}

func GenYamlTree(cmd *cobra.Command, dir string) error { _ = "STUB: not implemented"; return nil }

func GenYamlTreeCustom(cmd *cobra.Command, dir string, filePrepender, linkHandler func(string) string) error {
	_ = "STUB: not implemented"
	return nil
}

func GenYaml(cmd *cobra.Command, w io.Writer) error { _ = "STUB: not implemented"; return nil }

func GenYamlCustom(cmd *cobra.Command, w io.Writer, linkHandler func(string) string) error {
	_ = "STUB: not implemented"
	return nil
}

func genFlagResult(flags *pflag.FlagSet) []cmdOption { _ = "STUB: not implemented"; return nil }
