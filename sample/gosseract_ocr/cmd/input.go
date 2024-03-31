package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Add input file path",
	Long:  `OCRX input command is used to provide the path of the target image`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Input file path logic
		fmt.Printf("The input command is called: %s\n", args[0])
	},
	Example: `ocrx input your/file/path.img`,
}

func init() {
	rootCmd.AddCommand(inputCmd)
}
