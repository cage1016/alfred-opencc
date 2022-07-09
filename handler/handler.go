package handler

import (
	"log"
	"sort"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"

	"github.com/cage1016/alfred-opencc/occ"
)

func OpenChineseConvertHandler(wf *aw.Workflow, lc map[occ.Item]occ.ConverMap, args []string) error {
	os := make([]occ.Item, 0, len(lc))

	for item := range lc {
		os = append(os, item)
	}

	sort.Sort(occ.Items(os))
	for _, k := range os {
		cm := lc[k]
		out, err := cm.Cc.Convert(strings.Join(args, " "))
		if err != nil {
			log.Printf("%s", err)
			continue
		}

		wf.NewItem(out).
			Subtitle(cm.Subtitle).
			Icon(&aw.Icon{Value: cm.Icon}).
			Arg(out).
			UID(strconv.Itoa(k.Order)).
			Valid(true)
	}

	return nil
}
