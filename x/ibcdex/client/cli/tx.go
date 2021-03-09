package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdCreateSlog())
	cmd.AddCommand(CmdUpdateSlog())
	cmd.AddCommand(CmdDeleteSlog())

	cmd.AddCommand(CmdCancelBuyOrder())

	cmd.AddCommand(CmdCancelSellOrder())

	cmd.AddCommand(CmdSendBuyOrder())

	cmd.AddCommand(CmdSendSellOrder())

	cmd.AddCommand(CmdSendCreatePair())

	return cmd
}
