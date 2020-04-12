package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func PrintTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "erik",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey, the current time is ", prettyTime)
			return nil
		},
	}
}
