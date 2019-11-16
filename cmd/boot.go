package cmd

import "github.com/spf13/cobra"

func init() {
	BootCMD.AddCommand(showVersionCMD)

	initFlags()
}

func initFlags()  {
	BootCMD.PersistentFlags().BoolP(showVersionCMD.Use, showVersionCMD.Use[:1], false, showVersionCMD.Short)
}

const Version  = "0.1.0"

var BootCMD = &cobra.Command{
	Use:   "SQL-Image",
	Short: "SQL docker-images!",
	Args: cobra.MaximumNArgs(1),
	Run:  doSQL,
}

var showVersionCMD = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run:   showVersion,
}