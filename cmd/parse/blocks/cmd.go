package blocks

import (
	"github.com/spf13/cobra"

	parsecmdtypes "github.com/saifullah619/juno/v3/cmd/parse/types"
)

// NewBlocksCmd returns the Cobra command that allows to fix all the things related to blocks
func NewBlocksCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks",
		Short: "Fix things related to blocks and transactions",
	}

	cmd.AddCommand(
		newAllCmd(parseConfig),
	)

	return cmd
}
