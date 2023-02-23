package demo

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:  "demo",
	Long: "this is a demo cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
