package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thinkgos/go-core-package/builder"

	"github.com/thinkgos/npsocks/pkg/tip"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Get version info",
	Example: fmt.Sprintf("%s version", builder.Name),
	RunE: func(*cobra.Command, []string) error {
		tip.PrintVersion()
		return nil
	},
}
