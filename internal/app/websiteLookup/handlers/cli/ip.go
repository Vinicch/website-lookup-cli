package cli

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Looks up the IP addresses for a particular host",
	RunE:  lookupIP,
}

func init() {
	ipCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(ipCmd)
}

func lookupIP(cmd *cobra.Command, args []string) error {
	ips, err := net.LookupIP(address)
	if err != nil {
		return err
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}
