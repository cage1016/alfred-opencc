package handler

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	aw "github.com/deanishe/awgo"

	"github.com/cage1016/alfred-opencc/occ"
)

func OpenChineseConvertHandler(wf *aw.Workflow, lc map[occ.Item]occ.ConvertMap, args []string) error {
	os := make([]occ.Item, 0, len(lc))
	buf := map[int]string{}

	for item := range lc {
		os = append(os, item)
	}

	sort.Sort(occ.Items(os))
	for _, v := range os {
		cm := lc[v]
		out, err := cm.Cc.Convert(args[0])
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		buf[v.Order] = out

		wf.NewItem(buf[v.Order]).
			Subtitle(fmt.Sprintf("⌘+L, ↩ Copy %s", cm.Subtitle)).
			Icon(&aw.Icon{Value: cm.Icon}).
			Arg(buf[v.Order]).
			Largetype(buf[v.Order]).
			UID(strconv.Itoa(v.Order)).
			Valid(true)
	}

	return nil
}
