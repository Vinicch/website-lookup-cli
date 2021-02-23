package cli

import (
	"fmt"
	"net"
	"websiteLookup/internal/app/websiteLookup/global"

	"github.com/spf13/cobra"
)

func showVersion(cmd *cobra.Command, args []string) {
	if version {
		fmt.Printf("Website Lookup CLI v%v\n", global.Version)
	} else {
		cmd.Help()
	}
}

func lookupCNAME(cmd *cobra.Command, args []string) error {
	commonName, err := net.LookupCNAME(cmd.Flag("address").Value.String())
	if err != nil {
		return err
	}
	fmt.Println(commonName)
	return nil
}

func lookupIP(cmd *cobra.Command, args []string) error {
	ips, err := net.LookupIP(cmd.Flag("address").Value.String())
	if err != nil {
		return err
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func lookupMX(cmd *cobra.Command, args []string) error {
	mxRecords, err := net.LookupMX(cmd.Flag("address").Value.String())
	if err != nil {
		return err
	}
	for _, record := range mxRecords {
		fmt.Println(record.Host, record.Pref)
	}
	return nil
}

func lookupNS(cmd *cobra.Command, args []string) error {
	nameServers, err := net.LookupNS(cmd.Flag("address").Value.String())
	if err != nil {
		return err
	}
	for _, nameServer := range nameServers {
		fmt.Println(nameServer.Host)
	}
	return nil
}
