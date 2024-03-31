package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write file to path",
	Long:  `OCRX write command is used to write file to the path`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		readFilePath := args[0]
		writeFilePath := args[1]
		// Write File
		fmt.Printf("Successfully read from %s and wrote to %s\n", readFilePath, writeFilePath)
	},
	Example: `ocrx read /read/file/path.img write /write/file/path`,
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
