package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"heya/x/tokenfactory/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Token factory transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		NewCreateDenomCmd(),
		NewMintCmd(),
		NewBurnCmd(),
		NewChangeAdminCmd(),
		NewForceTransferCmd(),
	)
	return txCmd
}

func NewCreateDenomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-denom [subdenom]",
		Short:   "Create a new denom",
		Long:    "Create a new denom. The sender becomes the admin. Denom format: factory/{creator}/{subdenom}",
		Example: version.AppName + " tx tokenfactory create-denom mytoken",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateDenom(clientCtx.GetFromAddress().String(), args[0])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint [amount]",
		Short:   "Mint tokens for a denom you are admin of",
		Long:    "Mint tokens for a denom. The amount must include the full denom (e.g. 1000factory/heya1.../mytoken). Use --mint-to to send to a different address.",
		Example: version.AppName + " tx tokenfactory mint 1000000factory/heya1abc123/mytoken",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			mintTo, _ := cmd.Flags().GetString("mint-to")
			msg := types.NewMsgMint(clientCtx.GetFromAddress().String(), args[0], mintTo)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String("mint-to", "", "address to mint tokens to (defaults to sender)")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBurnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn [amount]",
		Short:   "Burn tokens for a denom you are admin of",
		Long:    "Burn tokens. The amount must include the full denom (e.g. 1000factory/heya1.../mytoken).",
		Example: version.AppName + " tx tokenfactory burn 1000000factory/heya1abc123/mytoken",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgBurn(clientCtx.GetFromAddress().String(), args[0])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewChangeAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "change-admin [denom] [new-admin]",
		Short:   "Change the admin of a denom",
		Long:    "Change the admin of a denom. Only the current admin can execute this.",
		Example: version.AppName + " tx tokenfactory change-admin factory/heya1abc123/mytoken heya1xyz789",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgChangeAdmin(clientCtx.GetFromAddress().String(), args[0], args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewForceTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "force-transfer [amount] [dest-addr]",
		Short:   "Force transfer tokens to an address (admin only)",
		Long:    "Forcefully transfer tokens from any holder to a destination address. Only the denom admin can execute this.",
		Example: version.AppName + " tx tokenfactory force-transfer 1000000factory/heya1abc123/mytoken heya1destaddr",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgForceTransfer(clientCtx.GetFromAddress().String(), args[0], args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
