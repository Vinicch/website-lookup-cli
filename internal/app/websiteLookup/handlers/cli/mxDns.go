package cli

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var mxCmd = &cobra.Command{
	Use:   "mx",
	Short: "Looks up the Common Name for a particular host",
	RunE:  lookupMX,
}

func init() {
	mxCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(mxCmd)
}

func lookupMX(cmd *cobra.Command, args []string) error {
	mxRecords, err := net.LookupMX(address)
	if err != nil {
		return err
	}
	for _, record := range mxRecords {
		fmt.Println(record.Host, record.Pref)
	}
	return nil
}
