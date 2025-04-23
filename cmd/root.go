package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "my-cli",
    Short: "My CLI is a CLI tool that serve absolutely nothing",
    Long: "My CLI is a CLI tool that serve absolutely nothing, and it is useless as its creator",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello world!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

