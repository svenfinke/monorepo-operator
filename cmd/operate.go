package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// operateCmd represents the operate command
var operateCmd = &cobra.Command{
	Use:   "operate",
	Short: "Execute a command on your git repositories",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("operate called")
	},
}

func init() {
	rootCmd.AddCommand(operateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// operateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// operateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}