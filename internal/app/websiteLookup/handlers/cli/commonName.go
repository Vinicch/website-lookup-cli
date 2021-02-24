package cli

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var cnCmd = &cobra.Command{
	Use:   "cn",
	Short: "Looks up the Common Name for a particular host",
	RunE:  lookupCNAME,
}

func init() {
	cnCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(cnCmd)
}

func lookupCNAME(cmd *cobra.Command, args []string) error {
	commonName, err := net.LookupCNAME(address)
	if err != nil {
		return err
	}
	fmt.Println(commonName)
	return nil
}
