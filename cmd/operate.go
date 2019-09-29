package cmd

import (
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"strings"
)

// operateCmd represents the operate command
var operateCmd = &cobra.Command{
	Use:   "operate",
	Short: "Execute a command on your git repositories",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(".mono-operator.yml")
		if err != nil {
			panic(err.Error())
		}

		m.Exec(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(operateCmd)
}
