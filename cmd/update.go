/*
Copyright © 2022 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Open Chinese Convert",
	Run: func(cmd *cobra.Command, args []string) {
		wf.Configure(aw.TextErrors(true))
		log.Println("Checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}