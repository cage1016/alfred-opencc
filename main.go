package main

import (
	"errors"
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"

	"github.com/cage1016/alfred-opencc/alfred"
	"github.com/cage1016/alfred-opencc/handler"
	"github.com/cage1016/alfred-opencc/occ"
)

const updateJobName = "checkForUpdate"

var (
	repo = "cage1016/alfred-opencc"
	wf   *aw.Workflow
)

var cfg occ.Config

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	cfg.S2t_Enabled = alfred.GetENABLED_S2T(wf)
	cfg.T2s_Enabled = alfred.GetENABLED_T2S(wf)
	cfg.S2tw_Enabled = alfred.GetENABLED_S2TW(wf)
	cfg.Tw2s_Enabled = alfred.GetENABLED_TW2S(wf)
	cfg.S2hk_Enabled = alfred.GetENABLED_S2HK(wf)
	cfg.Hk2s_Enabled = alfred.GetENABLED_HK2S(wf)
	cfg.S2twp_Enabled = alfred.GetENABLED_S2TWP(wf)
	cfg.Tw2sp_Enabled = alfred.GetENABLED_TW2SP(wf)
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.FatalError(errors.New("please provide some input 👀"))
	}

	handlers := map[string]func(*aw.Workflow, map[occ.Item]occ.ConverMap, []string) error{
		"occ": handler.OpenChineseConvertHandler,
		"update": func(wf *aw.Workflow, _ map[occ.Item]occ.ConverMap, _ []string) error {
			wf.Configure(aw.TextErrors(true))
			log.Println("Checking for updates...")
			if err := wf.CheckForUpdate(); err != nil {
				wf.FatalError(err)
			}
			return nil
		},
	}

	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "update")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(&aw.Icon{Value: "update-available.png"})
	}

	h, found := handlers[args[0]]
	if !found {
		wf.FatalError(errors.New("command not recognized 👀"))
	}

	err := h(wf, occ.New(cfg), args[1:])
	if err != nil {
		wf.FatalError(err)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
