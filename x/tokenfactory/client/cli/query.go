package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"heya/x/tokenfactory/types"
)

func NewQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Token factory query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		NewDenomAdminCmd(),
		NewParamsCmd(),
	)
	return queryCmd
}

func NewDenomAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "denom-admin [denom]",
		Short:   "Get the admin of a denom",
		Args:    cobra.ExactArgs(1),
		Example: "heyad query tokenfactory denom-admin factory/heya1abc123/mytoken",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			req := &types.QueryDenomAdminRequest{Denom: args[0]}
			resp, err := types.NewQueryClient(clientCtx).DenomAdmin(cmd.Context(), req)
			if err != nil {
				return err
			}
			bz, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			return clientCtx.PrintString(string(bz) + "\n")
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Get token factory parameters",
		Args:    cobra.NoArgs,
		Example: "heyad query tokenfactory params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			req := &types.QueryParamsRequest{}
			resp, err := types.NewQueryClient(clientCtx).Params(cmd.Context(), req)
			if err != nil {
				return err
			}
			bz, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			return clientCtx.PrintString(string(bz) + "\n")
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
