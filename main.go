package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vSivarajah/AirlineReservation/app"
	"github.com/vSivarajah/AirlineReservation/cmd"
)

func main() {

	rootCmd := &cobra.Command{
		Use:          "test",
		Short:        "Hello World!",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(cmd.PrintTimeCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	app.StartApplication()
}
