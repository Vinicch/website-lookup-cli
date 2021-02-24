package cli

import (
	"fmt"
	"websiteLookup/internal/app/websiteLookup/global"

	"github.com/spf13/cobra"
)

var (
	version bool
	address string
	rootCmd = &cobra.Command{
		Use:   "wlc",
		Short: "Let's you query IPs, CNAMEs, MX records and Name Servers!",
		Run:   showVersion,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Gets the application version")
}

func showVersion(cmd *cobra.Command, args []string) {
	if version {
		fmt.Printf("Website Lookup CLI v%v\n", global.Version)
	} else {
		cmd.Help()
	}
}

// Execute will initiate the application
func Execute() error {
	return rootCmd.Execute()
}
