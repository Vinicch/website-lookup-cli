package cli

var (
	version bool
	address string
)

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Gets the application version")

	cnCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	ipCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	mxCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	nsCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")

	rootCmd.AddCommand(cnCmd)
	rootCmd.AddCommand(ipCmd)
	rootCmd.AddCommand(mxCmd)
	rootCmd.AddCommand(nsCmd)
}

// Execute will initiate the application
func Execute() error {
	return rootCmd.Execute()
}
