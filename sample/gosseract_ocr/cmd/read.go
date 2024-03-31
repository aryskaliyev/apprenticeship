package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"ocrx/pkg"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read file from path",
	Long:  `OCRX read command is used to read image from the path`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.ConvertImageToText(args[0], "")
		if err != nil {
			color.Red(err.Error())
			return
		}
		color.Green("Success!")
	},
	Example: `ocrx read your/file/path.img`,
}

func init() {
	rootCmd.AddCommand(readCmd)
}
