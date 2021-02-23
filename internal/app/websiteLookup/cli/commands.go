package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wlc",
		Short: "Let's you query IPs, CNAMEs, MX records and Name Servers!",
		Run:   showVersion,
	}
	cnCmd = &cobra.Command{
		Use:   "cn",
		Short: "Looks up the Common Name for a particular host",
		RunE:  lookupCNAME,
	}
	ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "Looks up the IP addresses for a particular host",
		RunE:  lookupIP,
	}
	mxCmd = &cobra.Command{
		Use:   "mx",
		Short: "Looks up the Common Name for a particular host",
		RunE:  lookupMX,
	}
	nsCmd = &cobra.Command{
		Use:   "ns",
		Short: "Looks up the Name Servers for a particular host",
		RunE:  lookupNS,
	}
)
