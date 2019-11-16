package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func doSQL(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		if showVersionFlag, _ := cmd.PersistentFlags().GetBool(showVersionCMD.Use); showVersionFlag {
			showVersion(cmd, args)
			return
		}

		// todo
		fmt.Println("into real!")
	} else {
		imageTag := args[0]
		// todo
		fmt.Println(imageTag)

		os.Exit(1)
	}

}
