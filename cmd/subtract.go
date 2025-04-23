package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"subtraction"},
	Short:   "Subtract 2 numbers",
	Long:    "Carry out subtraction on 2 numbers specifically",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s and %s = %s. \n\n", args[0], args[1], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
