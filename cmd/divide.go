package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"div"},
	Short:   "Divide one number by another",
	Long:    "Carry out division on 2 numbers",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := Divide(args[0], args[1], roundUp)
		if err != nil {
			return err
		}
		fmt.Printf("Division of %s by %s = %s. \n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	divideCmd.Flags().BoolVarP(&roundUp, "round", "r", false, "Round results up to 2 decimal")
	rootCmd.AddCommand(divideCmd)
}
