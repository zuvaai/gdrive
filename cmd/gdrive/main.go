package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"gdrive/pkg/download"
)

func main() {
	rootCmd := getCmd()
	err := rootCmd.Execute()
	if err != nil {
		log.Printf("cmd error = %v", err)
	}
}

func getCmd() *cobra.Command {
	appName := filepath.Base(os.Args[0])
	cmd := &cobra.Command{
		Use: appName,
	}
	cmd.AddCommand(download.GetCmd())
	return cmd
}
