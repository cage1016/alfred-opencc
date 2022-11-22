/*
Copyright © 2022 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-opencc/alfred"
	"github.com/cage1016/alfred-opencc/occ"
)

// qoccCmd represents the qocc command
var qoccCmd = &cobra.Command{
	Use:   "qocc",
	Short: "Quick Convert Chinese between Traditional and Simplified",
	Run:   runQOCCCmd,
}

func runQOCCCmd(c *cobra.Command, args []string) {
	lc := occ.New(cfg)

	x := lc[alfred.GetQUICK_OCC(wf)]
	if !x.Enabled {
		ErrorHandle(x.Subtitle, "末啟用，請至 Configure workflow/Quikc OCC 開啟")
		return
	}
	out, _ := x.Cc.Convert(args[0])

	av.Arg(out)
	av.Var("title", x.Subtitle)
	av.Var("message", out)
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(qoccCmd)
}
