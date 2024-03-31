package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"ocrx/pkg"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write file to path",
	Long:  `OCRX write command is used to write file to the path`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.ConvertImageToText(args[0], args[1])
		if err != nil {
			color.Red(err.Error())
			return
		}
		color.Green("Success!")
	},
	Example: `ocrx read /read/file/path.img write /write/file/path`,
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
