/*
Copyright © 2022 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-opencc/occ"
)

// occCmd represents the occ command
var occCmd = &cobra.Command{
	Use:   "occ",
	Short: "Convert Chinese between Traditional and Simplified",
	Run:   runOCCCmd,
}

func runOCCCmd(c *cobra.Command, args []string) {
	lc := occ.New(cfg)

	buf := map[int]string{}
	os := []occ.ConvertMap{}
	for _, v := range lc {
		if v.Enabled {
			os = append(os, v)
		}
	}
	sort.SliceStable(os, func(i, j int) bool {
		return os[i].Order < os[j].Order
	})

	wf.Configure(aw.SuppressUIDs(true))
	for _, v := range os {
		out, err := v.Cc.Convert(args[0])
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		buf[v.Order] = out

		wf.NewItem(buf[v.Order]).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %s", v.Subtitle)).
			Arg(buf[v.Order]).
			Valid(true).
			Largetype(buf[v.Order*-1]).
			UID(strconv.Itoa(v.Order)).
			Icon(&aw.Icon{Value: v.Icon}).
			Var("title", v.Subtitle).
			Var("message", buf[v.Order])
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(occCmd)
}
