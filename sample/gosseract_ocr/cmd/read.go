package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read file from path",
	Long:  `OCRX read command is used to read image from the path`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readFilePath := args[0]
		img, err := utils.ReadImage(readFilePath)
		if err != nil {
			fmt.Errorf(color(err))
			return
		}
		fmt.Printf("The read command is called: %s\n", args[0])
	},
	Example: `ocrx read your/file/path.img`,
}

func init() {
	rootCmd.AddCommand(readCmd)
}
