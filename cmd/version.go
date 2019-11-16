package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("SQL-IMAGE %s\n", Version)
}
