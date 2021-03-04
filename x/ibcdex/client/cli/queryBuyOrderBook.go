package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func CmdListBuyOrderBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-buyOrderBook",
		Short: "list all buyOrderBook",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBuyOrderBookRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BuyOrderBookAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowBuyOrderBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-buyOrderBook [index]",
		Short: "shows a buyOrderBook",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBuyOrderBookRequest{
				Index: args[0],
			}

			res, err := queryClient.BuyOrderBook(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
