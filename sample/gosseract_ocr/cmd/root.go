package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ocrx",
	Short: "A CLI tool to manage ocr",
	Long:  `OCRX is a CLI tool that offers a set of commands for managing text recognition from images with ease!`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {}
