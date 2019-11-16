package main

import (
	"fmt"
	"os"

	"github.com/KevinSmile/SQL-Image/cmd"
)

func main() {
	if err := cmd.BootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
