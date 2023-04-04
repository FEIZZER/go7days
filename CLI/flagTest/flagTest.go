package flagTest

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:  "flagCmdRoot",
	Long: "this is a flagTest cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("filePath:%+v, model:%+v, count:%+v \n", filePath, model, count)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var (
	filePath string
	model    string
	count    int32
)

func init() {
	filePath = *rootCmd.PersistentFlags().StringP("filePath", "f", "default_path", "usage")
	rootCmd.PersistentFlags().StringVar(&model, "model", "default_model", "usage")
	rootCmd.PersistentFlags().Int32VarP(&count, "count", "c", 1, "usage")
}
