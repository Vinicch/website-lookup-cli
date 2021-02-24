package cli

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Looks up the Name Servers for a particular host",
	RunE:  lookupNS,
}

func init() {
	nsCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(nsCmd)
}

func lookupNS(cmd *cobra.Command, args []string) error {
	nameServers, err := net.LookupNS(address)
	if err != nil {
		return err
	}
	for _, nameServer := range nameServers {
		fmt.Println(nameServer.Host)
	}
	return nil
}
