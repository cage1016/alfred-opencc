/*
Copyright © 2022 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-opencc/alfred"
	"github.com/cage1016/alfred-opencc/occ"
)

const updateJobName = "checkForUpdate"

var (
	repo = "cage1016/alfred-opencc"
	wf   *aw.Workflow
	cfg  occ.Config
	av   = aw.NewArgVars()
)

func ErrorHandle(title, message string) {
	av.Var("title", title)
	av.Var("message", fmt.Sprintf("Error: %s", message))
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

func CheckForUpdate() {
	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")
		cmd := exec.Command(os.Args[0], "update")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(&aw.Icon{Value: "update-available.png"})
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Open Chinese Convert",
	Short: "Convert Chinese between Traditional and Simplified by OpenCC",
	Run: func(cmd *cobra.Command, args []string) {
		CheckForUpdate()
		wf.SendFeedback()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	wf.Run(func() {
		cfg.S2t_Enabled = alfred.GetENABLED_S2T(wf)
		cfg.T2s_Enabled = alfred.GetENABLED_T2S(wf)
		cfg.S2tw_Enabled = alfred.GetENABLED_S2TW(wf)
		cfg.Tw2s_Enabled = alfred.GetENABLED_TW2S(wf)
		cfg.S2hk_Enabled = alfred.GetENABLED_S2HK(wf)
		cfg.Hk2s_Enabled = alfred.GetENABLED_HK2S(wf)
		cfg.S2twp_Enabled = alfred.GetENABLED_S2TWP(wf)
		cfg.Tw2sp_Enabled = alfred.GetENABLED_TW2SP(wf)

		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	})
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	wf.Args() // magic for "workflow:update"
}
