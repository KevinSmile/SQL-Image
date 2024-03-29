package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/KevinSmile/SQL-Image/fsql-img/terminal"
	"github.com/spf13/cobra"
)

func doSQL(cmd *cobra.Command, args []string) {
	if showVersionFlag, _ := cmd.PersistentFlags().GetBool(showVersionCMD.Use); showVersionFlag {
		showVersion(cmd, args)
		return
	}

	fmt.Println("SQL docker-images: image -> database, layer-diff -> table")
	if err := terminal.Start(); err != nil {
		log.Fatal(err.Error())
	}
	os.Exit(0)
}
