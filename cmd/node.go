package cmd

import (
	"github.com/mentos1386/ipfs-cloud/pkg/ipfs"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:  "node",
	Short: "Start IPFS Node",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ipfs.StartNode()
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}