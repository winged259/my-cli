package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var roundUp bool

var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mul", "multiple", "multi"},
	Short:   "Multiply 2 numbers",
	Long:    "Carry out multiplication operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplication of %s and %s is %s \n\n", args[0], args[1], Multiply(args[0], args[1], roundUp))
	},
}

func init() {
	multiplyCmd.Flags().BoolVarP(&roundUp, "round", "r", false, "Round results up to 2 decimals")
	rootCmd.AddCommand(multiplyCmd)
}
